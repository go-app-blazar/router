package router

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/go-app-blazar/router/route"
	"github.com/maxence-charriere/go-app/v11/pkg/app"
)

// LayoutWrapper is a wrapper around the desired component that handles top-level events to ensure that the component is properly updated.
type LayoutWrapper struct {
	app.Compo

	// These need to be public so that the component is properly re-rendered.
	LayoutComponent        app.Composer
	Meta                   map[string]string // TODO: Consider using `any` and having wrapper functions to get the values.
	Route                  route.Route
	RouteVariables         map[string]string
	PathVariablesFunctions []func(ctx app.Context, variables map[string]string)
}

// TODO: REMOVE THIS IF WE DON'T NEED IT
func (c *LayoutWrapper) OnMount(ctx app.Context) {
	slog.InfoContext(ctx.Context, "LayoutWrapper: OnMount", "url", ctx.Page().URL())
}

func (c *LayoutWrapper) OnNav(ctx app.Context) {
	slog.InfoContext(ctx.Context, "LayoutWrapper: OnNav", "url", ctx.Page().URL())
	slog.InfoContext(ctx.Context, "LayoutWrapper: OnNav", "self", fmt.Sprintf("%p", c), "LayoutComponent", fmt.Sprintf("%T", c.LayoutComponent), "LayoutComponentPointer", fmt.Sprintf("%p", c.LayoutComponent))

	matched, variables := c.Route.Match(ctx.Page().URL().Path)

	slog.InfoContext(ctx.Context, "LayoutWrapper: OnNav", "matched", matched)
	slog.InfoContext(ctx.Context, "LayoutWrapper: OnNav", "variables", variables)
	slog.InfoContext(ctx.Context, "LayoutWrapper: OnNav", "Component", fmt.Sprintf("%T", c.LayoutComponent))

	if !matched {
		slog.WarnContext(ctx.Context, "LayoutWrapper: OnNav: Could not match route somehow.", "route", c.Route, "path", ctx.Page().URL().Path)
		return
	}

	c.RouteVariables = variables

	for _, f := range c.PathVariablesFunctions {
		if f == nil {
			continue
		}
		f(ctx, c.RouteVariables)
	}

	activeRoute := ActiveRoute{
		Path:      ctx.Page().URL().Path,
		Meta:      c.Meta,
		Variables: c.RouteVariables,
	}
	slog.InfoContext(ctx.Context, "LayoutWrapper: OnNav: Setting active route.", "activeRoute", activeRoute)
	ctx.SetState(StateRoute, activeRoute)
}

// TODO: REMOVE THIS IF WE DON'T NEED IT
func (c *LayoutWrapper) OnUpdate(ctx app.Context) {
	slog.InfoContext(ctx.Context, "LayoutWrapper: OnUpdate", "url", ctx.Page().URL())
}

// TODO: It looks like we want to leverage OnMount (first time) and OnUpdate (subsequent times) to tell our components that something has changed.
// TODO: Or, consider adding a public Route property to the pages that need routes so that this info can automatically do what it needs to.

func (c *LayoutWrapper) Render() app.UI {
	slog.InfoContext(context.TODO(), "LayoutWrapper: Render")

	return c.LayoutComponent
}
