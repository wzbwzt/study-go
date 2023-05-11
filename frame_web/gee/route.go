package gee

import (
	"log"
	"net/http"
	"strings"
)

type router struct {
	roots    map[string]*node
	handlers map[string]HandlerFunc
}

// roots key eg, roots['GET'] roots['POST']
// handlers key eg, handlers['GET-/p/:lang/doc'], handlers['POST-/p/book']

func newRouter() *router {
	return &router{
		handlers: make(map[string]HandlerFunc),
		roots:    make(map[string]*node),
	}
}

// Only one * is allowed
func parsePattern(pattern string) []string {
	ps := strings.Split(pattern, "/")
	out := make([]string, 0)
	for _, v := range ps {
		if v != "" {
			out = append(out, v)
			if strings.HasPrefix(v, "*") {
				break
			}
		}
	}
	return out
}

func (r *router) addRoute(method string, pattern string, handler HandlerFunc) {
	log.Printf("Route %4s - %s", method, pattern)
	parts := parsePattern(pattern)
	_, ok := r.roots[method]
	if !ok {
		r.roots[method] = &node{}
	}
	r.roots[method].insert(pattern, parts, 0)
	key := method + "-" + pattern
	r.handlers[key] = handler
}

func (r *router) getRoute(method string, path string) (*node, map[string]string) {
	searchpaths := parsePattern(path)
	n, ok := r.roots[method]
	if !ok {
		return nil, nil
	}
	outmap := make(map[string]string)
	out := n.search(searchpaths, 0)
	if out != nil {
		ps := parsePattern(out.pattern)
		for index, part := range ps {
			if part[0] == ':' {
				outmap[part[1:]] = searchpaths[index]
			}
			if part[0] == '*' && len(part) > 1 {
				outmap[part[1:]] = strings.Join(searchpaths[index:], "/")
				break
			}
		}
		return out, outmap
	}
	return nil, nil
}

func (r *router) handle(c *Context) {
	node, param_map := r.getRoute(c.Method, c.Path)
	if node != nil {
		c.Params = param_map
		key := c.Method + "-" + node.pattern
		c.handlers = append(c.handlers, r.handlers[key])
	} else {
		c.handlers = append(c.handlers, func(c *Context) {
			c.String(http.StatusNotFound, "404 NOT FOUND: %s\n", c.Path)
		})
	}
	c.Next()
}

type node struct {
	pattern  string  // 待匹配路由，例如 /p/:lang
	part     string  // 路由中的一部分，例如 :lang
	children []*node // 子节点，例如 [doc, tutorial, intro]
	isWild   bool    // 是否精确匹配，part 含有 : 或 * 时为true
}

// 第一个匹配成功的节点，用于插入
func (n *node) matchChild(part string) *node {
	for _, v := range n.children {
		if v.part == part || v.isWild {
			return v
		}
	}
	return nil
}

// 所有匹配成功的节点，用于查找
func (n *node) matchChildren(part string) []*node {
	out := make([]*node, 0)
	for _, v := range n.children {
		if v.part == part || v.isWild {
			out = append(out, v)
		}
	}
	return out
}

func (n *node) insert(pattern string, parts []string, height int) {
	if len(parts) == height {
		n.pattern = pattern
		return
	}
	child := n.matchChild(parts[height])
	if child == nil {
		child = &node{
			part:     parts[height],
			children: make([]*node, 0),
			isWild:   parts[height][0] == ':' || parts[height][0] == '*',
		}
		n.children = append(n.children, child)
	}
	child.insert(pattern, parts, height+1)
}

func (n *node) search(parts []string, height int) *node {
	if len(parts) == height || strings.HasPrefix(n.part, "*") {
		if n.pattern != "" {
			return n
		}
		return nil
	}
	childs := n.matchChildren(parts[height])
	if len(childs) == 0 {
		return nil
	}
	for _, child := range childs {
		res := child.search(parts, height+1)
		if res != nil {
			return res
		}
	}
	return nil
}
