package controller

import (
	"time"

	"github.com/hashicorp/hcat"
)

var _ hcat.Templater = (template)(nil)
var _ hcat.Renderer = (template)(nil)

// template describes the interface for hashicat's Template structure
// which implements the interfaces Templater and Renderer
// https://github.com/hashicorp/hcat
type template interface {
	Render(content []byte) (hcat.RenderResult, error)
	Execute(hcat.Recaller) (*hcat.ExecuteResult, error)
	ID() string
}

// resolver describes the interface for hashicat's Resolver structure
// which does implement any hashicat interface at time of writing
// https://github.com/hashicorp/hcat
type resolver interface {
	Run(tmpl hcat.Templater, w hcat.Watcherer) (hcat.ResolveEvent, error)
}

var _ hcat.Watcherer = (watcher)(nil)

// watcher describes the interface for hashicat's Watcher structure
// which implements the interface Watcherer
// https://github.com/hashicorp/hcat
type watcher interface {
	Wait(timeout time.Duration) error
	Add(d hcat.Dependency) bool
	Changed(tmplID string) bool
	Recall(id string) (interface{}, bool)
	Register(tmplID string, deps ...hcat.Dependency)
}
