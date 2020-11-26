// Copyright (c) 2020 DistAlchemist
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package mql

import (
	"log"

	"github.com/antlr/antlr4/runtime/Go/antlr"

	"github.com/DistAlchemist/Mongongo/mql/parser"
)

type mqlListener struct {
	*parser.BaseMqlListener
	root    *node
	curNode *node
}

type node struct {
	id       int
	text     string
	parent   *node
	children []*node
}

func (n *node) addChild(id int, text string) *node {
	child := &node{id: id, text: text, parent: n}
	n.children = append(n.children, child)
	return child
}

func (n *node) dfs() {
	log.Printf("enter id: %2v, text: %v\n", n.id, n.text)
	for _, c := range n.children {
		c.dfs()
	}
	log.Printf("exit  id: %2v, text: %v\n", n.id, n.text)
}

func (l *mqlListener) init() {
	l.root = &node{-1, "", nil, nil}
	l.curNode = l.root
}

// EnterEveryRule is called when any rule is entered.
func (l *mqlListener) EnterEveryRule(c antlr.ParserRuleContext) {
	log.Printf("Enter: GetRuleIndex: %v, ", c.GetRuleIndex())
	log.Printf("GetText %v\n", c.GetText())
	l.curNode = l.curNode.addChild(c.GetRuleIndex(), c.GetText())
}

// ExitEveryRule is called when any rule is exited.
func (l *mqlListener) ExitEveryRule(c antlr.ParserRuleContext) {
	log.Printf("Enter: GetRuleIndex: %v, ", c.GetRuleIndex())
	log.Printf("GetText %v\n", c.GetText())
	l.curNode = l.curNode.parent
}
