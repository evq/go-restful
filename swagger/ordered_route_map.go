package swagger

import "github.com/evq/go-restful"

type orderedRouteMap struct {
	elements map[string][]restful.Route
	keys     []string
}

func newOrderedRouteMap() *orderedRouteMap {
	return &orderedRouteMap{
		elements: map[string][]restful.Route{},
		keys:     []string{},
	}
}

func (o *orderedRouteMap) Add(key string, route restful.Route) {
	routes, ok := o.elements[key]
	if ok {
		routes = append(routes, route)
		o.elements[key] = routes
		return
	}
	o.elements[key] = []restful.Route{route}
	o.keys = append(o.keys, key)
}

func (o *orderedRouteMap) Do(block func(key string, routes []restful.Route)) {
	for _, k := range o.keys {
		block(k, o.elements[k])
	}
}
