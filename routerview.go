package router

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/maxence-charriere/go-app/v11/pkg/app"
)

// RouterViewInterface is the interface that must be implemented in order for a component to be
// properly used in a route as a layout.
//
// Instead of implementing this interface, a component should embed `RouterViewComponent`.
type RouterViewInterface interface {
	RouterView() app.Composer
	SetRouterView(app.Composer)
}

// RouterViewComponent can be embedded in a layout component.
//
// The router wil set the specific router view component with `SetComponent`, and the
// embedding component can call `RouterViewComponent` in its `Render` function to get the component
// that should be rendered in the router view.
type RouterViewComponent struct {
	IRouterViewComponent app.Composer
}

var _ RouterViewInterface = (*RouterViewComponent)(nil)
var _ app.Updater = (*RouterViewComponent)(nil)

// RouterView returns the router view comonent.  Put this where you want the route component
// to be rendered.
//
// In Vue, this would be the `<router-view>` component.
func (v *RouterViewComponent) RouterView() app.Composer {
	return v.IRouterViewComponent
}

// SetRouterView is used by the router to set the component that will be returned by `RouterView`.
//
// This should not be called by anything but the router.
func (v *RouterViewComponent) SetRouterView(component app.Composer) {
	v.IRouterViewComponent = component
}

func (v *RouterViewComponent) OnUpdate(ctx app.Context) {
	slog.DebugContext(context.TODO(), "RouterViewComponent: Update.", "component", fmt.Sprintf("%T", v.IRouterViewComponent))
}
