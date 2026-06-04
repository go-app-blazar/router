![GitHub tag (latest SemVer)](https://img.shields.io/github/v/tag/go-app-blazar/router?label=version&logo=version&sort=semver)
[![PkgGoDev](https://pkg.go.dev/badge/github.com/go-app-blazar/router)](https://pkg.go.dev/github.com/go-app-blazar/router)


# router
A Vue-like router for go-app

# Usage
Register your routes with `router.Register` instead of using go-app's `RegisterRoute` functions.

```
router.Register(ctx,
	router.Route{
		Path: "/",
		Component: func() app.Composer {
			return &layout.CenterLayout{}
		},
		Meta: map[string]string{
			"require-login": "false",
		},
		Children: []router.Route{
			{
				Path: "/login",
				Component: func() app.Composer {
					return &page.LoginPage{}
				},
			},
			{
				Path: "/signup",
				Component: func() app.Composer {
					return &page.SignupPage{}
				},
			},
		},
	},
)
```

If a parent route's `Component` function returns a component that implements `router.RouterViewInterface`, then that component will be instantiated and `SetRouterView` will be called with the child component (and so on).

A component that implements `router.RouterViewInterface` should embed `router.RouterViewComponent` and include `RouterViewComponent.RouterView()` somewhere in its `Render` function.

For example:

```
type CenterLayout struct {
	app.Compo
	router.RouterViewComponent
}

var _ router.RouterViewInterface = (*CenterLayout)(nil)

func (c *CenterLayout) Render() app.UI {
	return app.Div().
		Class("center-layout").
		Body(
			c.RouterViewComponent.RouterView(),
		)
}
```