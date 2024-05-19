package git
import  "github.com/go-git/go-git/v5"

type Repo struct {
  repo *git.Repository   
}


func NewRepo(path string) (*Repo, error) {
  r, err := git.PlainOpen(path)
  if err != nil {
    return nil, err
  }

  return &Repo{repo: r}, nil
}
