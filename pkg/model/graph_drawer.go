package model

import (
	"fmt"
	"gitty/pkg/git"
	"sort"

	"golang.org/x/exp/maps"
)
const DefaultCanvasWidth = 100
const NodeConnectorUp = '│'
const NodeConnecterHorizontal = '─'
const NodeConnectorBranch = '╯'
const Node = 'o'
const NodeConnecterHorizontalBranch = '├'

type GraphDrawer struct {
  *git.LogGraph
  Width int
  canvas []rune
  ptr int
}


func (g *GraphDrawer) String() string {
  canvas := make([]rune, 1000000) 
  g.canvas = canvas
  rootCommit := g.LogGraph.RootCommits[0]
  g.nextLine()
  g.addLine(0, 0)
  g.drawCommit(rootCommit, 0) 
  g.nextLine()
  
  g.buildCanvas(g.LogGraph.RootCommits[0], 0)
  fmt.Println(string(g.canvas))
  return string(canvas)
}

func (g *GraphDrawer) buildCanvas(commit *git.GraphCommit, indent int) {
  commits := maps.Values(commit.Children)
  sort.Slice(commits, func(i, j int) bool {
    return commits[i].Length < commits[j].Length
  })
  
  if len(commits) == 0 {
    return 
  }

  mainCommit := commits[len(commits) - 1]
  commits = commits[:len(commits) - 1]
  for _, commit := range commits {
    g.addLine(0, indent)
    g.nextLine()

    g.buildBranch(commit, indent) 
    g.buildCanvas(commit, indent + 2)
  }

  g.addLine(0, indent)
  g.nextLine()

  g.buildMainBranch(mainCommit, indent)
  g.buildCanvas(mainCommit, indent)
} 

func (g *GraphDrawer) buildMainBranch(commit *git.GraphCommit, indent int) {
  g.addLine(0, indent)
  g.addLine(1, indent)

  g.Write(NodeConnectorUp,0, indent)

  g.nextLine()
  g.drawCommit(commit, indent)
  g.nextLine()
}
func (g *GraphDrawer) buildBranch(commit *git.GraphCommit, indent int) {
  g.addLine(0, indent + 1)
  g.addLine(1, indent + 1)

  g.Write(NodeConnecterHorizontalBranch,0, indent)
  g.Write(NodeConnecterHorizontal,0, indent+1)
  g.Write(NodeConnectorBranch,0, indent+2)
  
  g.nextLine()
  g.drawCommit(commit, indent+2)
  g.nextLine()
}

func (g *GraphDrawer) Write(letter rune, i, j int) {
  g.canvas[len(g.canvas) - ((g.ptr + i)*g.Width) + j - 1] = letter 
}

func (g *GraphDrawer) drawCommit(commit *git.GraphCommit, indent int) {
  g.Write(Node, 0, indent)
  ptr := indent + 2
  for _, letter := range commit.Hash.String() {
    g.Write(letter,0, ptr)
    ptr++
  }
}

func (g *GraphDrawer) addLine(line, indent int) {
  for i := 0; i < g.Width; i++ {
    switch {
    case i % 2 == 0 && i <= indent:
      g.Write(NodeConnectorUp, line, i)
    case i + 1 == g.Width:
      g.Write('\n', line, i)
    default:
      g.Write(' ', line, i)
    }
  }
}

func (g *GraphDrawer) nextLine() {
  g.ptr += 1 
}





