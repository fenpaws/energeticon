package handlers

import (
	"github.com/go-mojito/mojito"
	"time"
)

func HomeHandler(ctx mojito.RendererContext, cache mojito.Cache) {
	lastVisit := time.Time{}
	cache.GetOrDefault("lastVisit", &lastVisit, time.Now())
	cache.Set("lastVisit", time.Now())

	ctx.ViewBag().Set("lastVisit", lastVisit)
	ctx.MustView("Home")
}
