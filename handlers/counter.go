package handlers

import (
	"net/http"
	"sync"

	"github.com/swarnimcodes/gopher-tools/components"
)

type CounterHandler struct {
	count int
	mu    sync.Mutex
}

func NewCounterHandler(initCount int) *CounterHandler {
	return &CounterHandler{
		count: initCount,
		mu:    sync.Mutex{},
	}
}

func (h *CounterHandler) Decrement(w http.ResponseWriter, r *http.Request) {
	h.mu.Lock()
	h.count--
	count := h.count
	h.mu.Unlock()

	component := components.CountDisplay(count)
	component.Render(r.Context(), w)

}

func (h *CounterHandler) Increment(w http.ResponseWriter, r *http.Request) {
	h.mu.Lock()
	h.count++
	count := h.count
	h.mu.Unlock()

	component := components.CountDisplay(count)
	component.Render(r.Context(), w)

}
