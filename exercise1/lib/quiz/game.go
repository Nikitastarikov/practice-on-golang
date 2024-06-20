package libquiz

import (
	"bufio"
	"encoding/csv"
	"os"
	"strconv"
	"strings"
	"time"

	pkglog "github.com/Nikitastarikov/practice-on-golang/pkg/log"
)

type QuizGame struct {
	l           pkglog.Logger
	tasks       []Task
	timeToThink time.Duration
}

func NewQuizGame(l pkglog.Logger, fileName string, timeToThink int) (*QuizGame, error) {
	game := &QuizGame{
		l:           l,
		timeToThink: time.Duration(timeToThink) * time.Second,
	}

	err := game.read(fileName)

	if err != nil {
		return nil, err
	}

	return game, nil
}

func (g *QuizGame) Run() error {
	l := g.l.Named("QuizGame")
	l.Infof("game start!\n")
	l.Infof("time to think: %v\n", g.timeToThink)

	count := 0

	done := make(chan string)

	go g.input(done)

	for i := 0; i < len(g.tasks); i++ {
		l.Infof("Attention question: %v", g.tasks[i].question)

		timer := time.NewTimer(g.timeToThink)

		timeIsOver, res, err := g.handleAnswer(i, done, timer.C)

		if err != nil {
			l.Errorf("handle answer: %v", err)
		}

		count = count + res

		switch {
		case res > 0 && !timeIsOver:
			l.Infof("right answer")

		case timeIsOver:
			l.Infof("time is over")

		default:
			l.Infof("wrong answer")
		}
	}

	l.Infof("game over!\n")
	l.Infof("your count: %v", count)

	return nil
}

func (g *QuizGame) handleAnswer(taskNum int, done <-chan string, timer <-chan time.Time) (bool, int, error) {
	for {
		select {
		case answer := <-done:
			answer = strings.Trim(strings.ToLower(answer), "\n")
			if strings.Compare(answer, g.tasks[taskNum].answer) == 0 {
				return false, g.tasks[taskNum].count, nil
			} else {
				return false, 0, nil
			}

		case <-timer:
			return true, 0, nil
		}
	}
}

func (g *QuizGame) input(inputCh chan<- string) {
	l := g.l.Named("input")

	for {
		in := bufio.NewReader(os.Stdin)
		result, err := in.ReadString('\n')

		if err != nil {
			l.Errorf("read string: %v", err)
			return
		}

		inputCh <- result
	}
}

func (g *QuizGame) read(name string) error {
	l := g.l.Named("read")

	file, err := os.Open(name)

	if err != nil {
		l.Errorf("open file: %v", name)
		return err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	reader.FieldsPerRecord = 3
	reader.Comment = '#'

	records, err := reader.ReadAll()

	if err != nil {
		l.Errorf("read all: %v", err)
		return err
	}

	g.tasks = make([]Task, 0, len(records))

	for _, line := range records {
		countLine := strings.TrimSpace(line[2])
		count, err := strconv.Atoi(countLine)

		if err != nil {
			l.Errorf("atoi: %v", err)
			return err
		}

		task := Task{
			question: strings.TrimSpace(line[0]),
			answer:   strings.TrimSpace(line[1]),
			count:    count,
		}

		g.tasks = append(g.tasks, task)
	}

	return nil
}
