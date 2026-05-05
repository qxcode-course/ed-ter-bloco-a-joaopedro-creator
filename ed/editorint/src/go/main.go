package main

import (
	"fmt"

	"github.com/gdamore/tcell/v2"
)

type Editor struct {
	texto   *List[*List[rune]]
	it_line *Node[*List[rune]]
	it_char *Node[rune]
	screen  tcell.Screen
	style   tcell.Style
}

func (e *Editor) InsertChar(r rune) {
	e.it_char = e.it_line.Value.Insert(e.it_char, r)
	e.it_char = e.it_char.Next()
}

func (e *Editor) KeyLeft() {
	if e.it_char == e.it_line.Value.Front() {
        
        if e.it_line != e.texto.Front() {
            e.it_line = e.it_line.Prev()       
            e.it_char = e.it_line.Value.End()  
        }

    } else {
        e.it_char = e.it_char.Prev()
    }
}

func (e *Editor) KeyEnter() {
	nova := NewList[rune]()

    for it := e.it_char; it != e.it_line.Value.End(); {
        charParaMover := it.Value
        nova.PushBack(charParaMover) 
        it = e.it_line.Value.Erase(it)
    }
    e.texto.Insert(e.it_line.Next(), nova)
    e.it_line = e.it_line.Next()
    e.it_char = e.it_line.Value.Front()
}

func (e *Editor) KeyRight() {
	if e.it_char == e.it_line.Value.End() {
        if e.it_line.Next() != e.texto.End() {
            e.it_line = e.it_line.Next()        
            e.it_char = e.it_line.Value.Front()  
        }
    	} else {
        	e.it_char = e.it_char.Next()
    }
}

func (e *Editor) KeyUp() {
	if e.it_line == e.texto.Front() {
        return
    }
    colunaAlvo := e.it_line.Value.IndexOf(e.it_char)

    e.it_line = e.it_line.Prev()

    e.it_char = e.it_line.Value.Front()
    for i := 0; i < colunaAlvo; i++ {
        if e.it_char == e.it_line.Value.End() {
            break
        }
        e.it_char = e.it_char.Next()
    }
}

func (e *Editor) KeyDown() {
    if e.it_line.Next() == e.texto.End() {
        return
    }

    colunaAlvo := e.it_line.Value.IndexOf(e.it_char)

    e.it_line = e.it_line.Next()

    e.it_char = e.it_line.Value.Front()
    for i := 0; i < colunaAlvo; i++ {
        if e.it_char == e.it_line.Value.End() {
            break
        }
        e.it_char = e.it_char.Next()
    }
}

func (e *Editor) KeyBackspace() {
	if e.it_char == e.it_line.Value.Front(){
		if e.it_line == e.texto.Front(){
			return
		}

		linhaAtual := e.it_line
		linhaAcima := e.it_line.Prev()

		posicaoDeUniao := linhaAcima.Value.End()

		for it := linhaAtual.Value.Front(); it != linhaAtual.Value.End();{
			charParaMover := it.Value
            linhaAcima.Value.Insert(linhaAcima.Value.End(), charParaMover)
            
            it = linhaAtual.Value.Erase(it)
		}
		e.it_line = linhaAcima
        e.it_char = posicaoDeUniao      
        e.texto.Erase(linhaAtual)
	}else{
		e.it_char = e.it_line.Value.Erase(e.it_char.Prev())
	}
}

func (e *Editor) KeyDelete() {
	if e.it_char == e.it_line.Value.End() {
        
        // Só podemos juntar se houver uma linha abaixo
        if e.it_line.Next() != e.texto.End() {
            linhaAtual := e.it_line
            proximaLinha := e.it_line.Next()

            for it := proximaLinha.Value.Front(); it != proximaLinha.Value.End(); {
                charParaMover := it.Value
                linhaAtual.Value.Insert(linhaAtual.Value.End(), charParaMover)
                it = proximaLinha.Value.Erase(it)
            }

           
            e.texto.Erase(proximaLinha)
            
        }

    } else {
        e.it_char = e.it_line.Value.Erase(e.it_char)
    }
}

func main() {
	// Texto inicial e posição do cursor
	editor := NewEditor()
	editor.Draw()
	editor.MainLoop()
	defer editor.screen.Fini() // Encerra a tela ao sair
}

func (e *Editor) MainLoop() {
	for {
		ev := e.screen.PollEvent()
		switch ev := ev.(type) {
		case *tcell.EventKey:
			switch ev.Key() {
			case tcell.KeyEsc, tcell.KeyCtrlC:
				return
			case tcell.KeyEnter:
				e.KeyEnter()
			case tcell.KeyLeft:
				e.KeyLeft()
			case tcell.KeyRight:
				e.KeyRight()
			case tcell.KeyUp:
				e.KeyUp()
			case tcell.KeyDown:
				e.KeyDown()
			case tcell.KeyBackspace, tcell.KeyBackspace2:
				e.KeyBackspace()
			case tcell.KeyDelete:
				e.KeyDelete()
			default:
				if ev.Rune() != 0 {
					e.InsertChar(ev.Rune())
				}
			}
			e.Draw()
		case *tcell.EventResize:
			e.screen.Sync()
			e.Draw()
		}
	}
}

func NewEditor() *Editor {
	e := &Editor{}
	// Inicializa a tela
	screen, err := tcell.NewScreen()
	if err != nil {
		fmt.Printf("erro ao criar a tela: %v", err)
	}
	if err := screen.Init(); err != nil {
		fmt.Printf("erro ao iniciar a tela: %v", err)
	}
	e.screen = screen
	e.texto = NewList[*List[rune]]()
	e.texto.PushBack(NewList[rune]())
	e.it_line = e.texto.Front()
	e.it_char = e.it_line.Value.Back()
	// Define o estilo do texto (branco com fundo preto)
	e.style = tcell.StyleDefault.Foreground(tcell.ColorWhite).Background(tcell.ColorBlack)

	// Limpa a tela e define o estilo base
	e.screen.SetStyle(e.style)
	e.screen.Clear()
	return e
}

func (e *Editor) Draw() {
	e.screen.Clear()
	x := 0
	y := 0
	for line := e.texto.Front(); line != e.texto.End(); line = line.Next() {
		for char := line.Value.Front(); ; char = char.Next() {
			data := char.Value
			if char == line.Value.End() {
				data = '⤶'
			}
			if data == ' ' {
				data = '·'
			}
			if char == e.it_char {
				e.screen.SetContent(x, y, data, nil, e.style.Reverse(true))
			} else {
				e.screen.SetContent(x, y, data, nil, e.style)
			}
			x++
			if char == line.Value.End() {
				break
			}
		}
		y++
		x = 0
	}
	e.screen.Show()
}