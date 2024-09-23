package bookmarks

import (
	"fmt"
	"go-base/consoleColors"
)

const formatString = "%-5s %-10s %-5s\n"
const formatInt = "%-5d %-10s %-5s\n"

type BookmarkType = map[string]string

var welcomeText = consoleColors.Colors().RedBold(` 
										█▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀█	
										█░░╦─╦╔╗╦─╔╗╔╗╔╦╗╔╗░░█	
										█░░║║║╠─║─║─║║║║║╠─░░█	
										█░░╚╩╝╚╝╚╝╚╝╚╝╩─╩╚╝░░█	
										█▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄█
`)

func Bookmarks() {
	var bookmarks = BookmarkType{
		"YouTube":   "https://www.youtube.com",
		"Google":    "https://www.google.com",
		"Instagram": "https://www.instagram.com",
		"VK":        "https://www.vk.com",
	}
	color := consoleColors.Colors()
	fmt.Print(welcomeText)

	var userOutput int8
Menu:
	for {
		fmt.Print(color.Yellow(`
	1. Посмотреть закладки 
	2. Добавить закладку
	3. Удалить закладку
	4. Выход
`))
		fmt.Println(color.Green("________________________________"))
		fmt.Scanf("%d", &userOutput)
		fmt.Println(color.Green("________________________________"))

		switch userOutput {
		case 1:
			printBookmarks(bookmarks)
			fmt.Println(color.YellowBoldUl("________________________________"))
		case 2:
			addBookmark(bookmarks)
			fmt.Println(color.YellowBoldUl("________________________________"))
		case 3:
			removeBookmark(bookmarks)
			fmt.Println(color.YellowBoldUl("________________________________"))
		default:
			break Menu
		}
	}
}

func removeBookmark(bookmarks BookmarkType) {
	var removeBookmarkName string
	fmt.Print("Введите название закладки которую вы хотите удалить: ")
	fmt.Scanf("%s", &removeBookmarkName)
	delete(bookmarks, removeBookmarkName)
}

func addBookmark(bookmarks BookmarkType) {
	var bookmarkName string
	var bookmarkUrl string
	fmt.Print("Введите название закладки: ")
	fmt.Scanf("%s", &bookmarkName)
	fmt.Print("Введите URL для закладки: ")
	fmt.Scanf("%s", &bookmarkUrl)
	bookmarks[bookmarkName] = bookmarkUrl
}

func printBookmarks(bookmarks BookmarkType) {
	if len(bookmarks) == 0 {
		fmt.Println("У вас пока нет закладок")
	}
	fmt.Printf(formatString, "ID", "Name", "URL")
	index := 1
	for bookmark := range bookmarks {

		fmt.Printf(formatInt, index, bookmark, bookmarks[bookmark])
		index++
	}
}

//может переделаю с использованием библиотеки gocui
//var menuItems = []string{"Посмотреть закладки", "Добавить закладку", "Удалить закладку", "Выход"}
//
//func layout(g *gocui.Gui) error {
//	maxX, maxY := g.Size()
//
//	if v, err := g.SetView("menu", -1, -1, int(0.08*float32(maxX)), maxY-3); err != nil {
//		if err != gocui.ErrUnknownView {
//			return err
//		}
//		v.Highlight = true
//		v.SelFgColor = gocui.ColorGreen
//		v.SelBgColor = gocui.ColorBlack
//		v.SetCursor(0, 0)
//
//		//fmt.Fprint(v, welcomeText)
//
//		for _, item := range menuItems {
//			fmt.Fprintln(v, item)
//		}
//	}
//
//	if v, err := g.SetView("view", int(0.08*float32(maxX)), -1, maxX, maxY-3); err != nil {
//		if err != gocui.ErrUnknownView {
//			return err
//		}
//		v.Wrap = true
//	}
//
//	_, err := g.SetCurrentView("menu")
//	return err
//}
//
//func quit(g *gocui.Gui, v *gocui.View) error {
//	return gocui.ErrQuit
//}
//
//func cursorDown(g *gocui.Gui, v *gocui.View) error {
//	cx, cy := v.Cursor()
//	if cy < len(menuItems)-1 {
//		if err := v.SetCursor(cx, cy+1); err != nil {
//			return err
//		}
//	}
//	return nil
//}
//
//func cursorUp(g *gocui.Gui, v *gocui.View) error {
//	cx, cy := v.Cursor()
//	if cy > 0 {
//		if err := v.SetCursor(cx, cy-1); err != nil {
//			return err
//		}
//	}
//	return nil
//}
//
//func selectItem(g *gocui.Gui, v *gocui.View) error {
//	_, cy := v.Cursor()
//	switch menuItems[cy] {
//	case "Выход":
//		return gocui.ErrQuit
//	case "Посмотреть закладки":
//		printBookmarks(g)
//	case "Добавить закладку":
//		addBookmarks(g)
//
//	default:
//		fmt.Printf("Selected: %s\n", menuItems[cy])
//	}
//	return nil
//}
//
//func main() {
//	g, err := gocui.NewGui(gocui.OutputNormal)
//	if err != nil {
//		log.Panicln(err)
//	}
//	defer g.Close()
//
//	g.SetManagerFunc(layout)
//
//	if err := g.SetKeybinding("menu", gocui.KeyArrowDown, gocui.ModNone, cursorDown); err != nil {
//		log.Panicln(err)
//	}
//	if err := g.SetKeybinding("menu", gocui.KeyArrowUp, gocui.ModNone, cursorUp); err != nil {
//		log.Panicln(err)
//	}
//	if err := g.SetKeybinding("menu", gocui.KeyEnter, gocui.ModNone, selectItem); err != nil {
//		log.Panicln(err)
//	}
//	if err := g.SetKeybinding("", gocui.KeyCtrlC, gocui.ModNone, quit); err != nil {
//		log.Panicln(err)
//	}
//
//	if err := g.MainLoop(); err != nil && err != gocui.ErrQuit {
//		log.Panicln(err)
//	}
//}
//func printBookmarks(g *gocui.Gui) {
//	g.SetCurrentView("view")
//	v, err := g.View("view")
//
//	if err != nil {
//		log.Panicln(err)
//	}
//
//	v.Clear()
//	for key, item := range bookmarks {
//		fmt.Fprintln(v, key, item)
//	}
//}
//
//func addBookmarks(g *gocui.Gui) {
//	var bookmarkName string
//	var bookmarkUrl string
//	g.SetCurrentView("view")
//	v, err := g.View("view")
//
//	if err != nil {
//		log.Panicln(err)
//	}
//
//	v.Clear()
//	fmt.Fprint(v, "Введите название закладки: ")
//	fmt.Fscanln(v, &bookmarkName)
//	fmt.Fprint(v, "Введите адрес: ")
//	fmt.Fscanln(v, &bookmarkUrl)
//}
