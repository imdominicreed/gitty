package model

import (
	"gitty/pkg/git"
	"sort"
	"strings"

	"github.com/charmbracelet/lipgloss"
	"golang.org/x/exp/maps"
)
const DefaultCanvasWidth = 200
const NodeConnectorUp = '│'
const NodeConnecterHorizontal = '─'
const NodeConnectorBranch = '╯'
const Node = 'o'
const NodeConnecterHorizontalBranch = '├'
const Head = '@'

var HashColor = lipgloss.NewStyle().Foreground(lipgloss.Color("11"))
var BranchColor = lipgloss.NewStyle().Foreground(lipgloss.Color("2"))
var HeadColor = lipgloss.NewStyle().Foreground(lipgloss.Color("5"))

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
  g.addLine(0, -1)
  g.addLine(1, -1)
  g.drawCommit(rootCommit, 0) 
  g.nextLine()
  g.nextLine()
  
  g.buildCanvas(g.LogGraph.RootCommits[0], 0)

  s := string(canvas[len(canvas)- (g.Width * (g.ptr-1)):])

  return s
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

  g.drawCommit(commit, indent)
  g.nextLine()
  g.nextLine()
}
func (g *GraphDrawer) buildBranch(commit *git.GraphCommit, indent int) {
  g.addLine(0, indent + 1)
  g.addLine(1, indent + 1)

  g.Write(NodeConnecterHorizontalBranch,0, indent)
  g.Write(NodeConnecterHorizontal,0, indent+1)
  g.Write(NodeConnectorBranch,0, indent+2)
  
  g.drawCommit(commit, indent+2)
  g.nextLine()
  g.nextLine()
}

func (g *GraphDrawer) Write(letter rune, i, j int) {
  g.canvas[len(g.canvas) - ((g.ptr + i)*g.Width) + j ] = letter 
}

func (g *GraphDrawer) drawCommit(commit *git.GraphCommit, indent int) {
  node := Node  
  if commit.Head {
    node = Head
  }
  g.Write(node, 1, indent)
  indent += 2
  ptr := g.WriteString(HashColor.Render(commit.Hash.String()[:9]) + " ", 1, indent)
  ptr = g.WriteString(commit.Author.Name + " ", 1, ptr)
  for _, b := range commit.BranchTips {
    ptr = g.WriteString(BranchColor.Render(b.Reference.Name().Short()) + " ", 1, ptr)
  }

  titleEnd:= strings.Index(commit.Message, "\n")
  title := commit.Message[:titleEnd]
  if commit.Head {
    title = HeadColor.Render(title)
  }
  g.WriteString(title, 0, indent)

}

func (g *GraphDrawer) WriteString(s string, line, indent int) int {
  ptr := indent
   for _, letter := range s {
     g.Write(letter, line, ptr)
     ptr += 1
  }
  return ptr
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





