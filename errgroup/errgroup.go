package errgroup

import (
	"context"
	"sync"

	"git.eth4.dev/golibs/errors"
)

// Group - набор горутин для запуска подзадач некоторой большой задачи
// с возможностью обработки ошибок каждой из подзадач и возврата наверх
// их списка в виде `errors.Chain`
type Group struct {
	ctx         context.Context
	err         error
	cancel      func()
	concurrence chan struct{}
	wg          sync.WaitGroup
	errLock     sync.Mutex
}

// New - конструктор пустой группы, юзкейс когда отрабатывают все подзадачи,
// невзирая на ошибки в некоторых из них или во всех
func New() *Group {
	return &Group{}
}

// WithCancelOnErr - конструктор группы с унаследованным контекстом, юзкейс в
// котором при возникновении первой ошибки в одной из подзадач - выполнение всех подзадач прерывается
func WithCancelOnErr(ctx context.Context) *Group {
	ctx, cancel := context.WithCancel(ctx)

	return &Group{
		cancel: cancel,
		ctx:    ctx,
	}
}

// WithMaxConcurrency - устанавливает максимальное количество выполняемых
// одновременно подзадач
func (g *Group) WithMaxConcurrency(factor int) *Group {
	g.concurrence = make(chan struct{}, factor)

	return g
}

// Context - возвращает родительский контекст.
func (g *Group) Context() context.Context {
	return g.ctx
}

// Wait - ожидание завершения работы подзадач. Возвращает первую ошибку в случае
// использования кейса прерывания по ошибке и набор ошибок объединенных в `errors.Chain` с
// помощью And в случае кейса непрерывной работы
func (g *Group) Wait() error {
	g.wg.Wait()

	if g.cancel != nil {
		g.cancel()
	}

	return g.err
}

// Go - запуск выполнения подзадачи в группу
func (g *Group) Go(f func() error) {
	g.wg.Add(1)
	g.acquire()

	go func() {
		defer g.wg.Done()
		defer g.release()

		if err := f(); err != nil {
			g.errLock.Lock()
			if g.cancel != nil {
				if g.err == nil {
					g.cancel()
				}

				if g.err == nil {
					g.err = err
				}
			} else {
				g.err = errors.And(g.err, err)
			}
			g.errLock.Unlock()
		}
	}()
}

func (g *Group) acquire() {
	if cap(g.concurrence) > 0 {
		g.concurrence <- struct{}{}
	}
}

func (g *Group) release() {
	if cap(g.concurrence) > 0 {
		<-g.concurrence
	}
}
