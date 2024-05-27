package main

import "fmt"

func main() {

	getUp()

	const foresterHouse int = 1
	const forestLake int = 2
	const enchantedTree int = 3
	const pocketItemToBeSmall int = 1
	const pocketItemToBeBig int = 2
	const pocketItemToBeGreen int = 3
	const pocketItemMatches int = 4

	fmt.Printf("\n Оберіть шлях. Варіанти: 1 - будиночок лісника; 2 - озеро; 3 - зачароване дерево.\n")
	var whereToGo int
	fmt.Scan(&whereToGo)

	var stopAsking bool
	// for stopasking != true  is the same as !stopAsking
	for !stopAsking {
		if whereToGo == foresterHouse {
			fmt.Println("---> Ви йдете до будиночку лісника")
			stopAsking = true
		} else if whereToGo == forestLake {
			fmt.Println("---> Ви йдете до озера")
			stopAsking = true
		} else if whereToGo == enchantedTree {
			fmt.Printf("---> Ви йдете до зачарованого дерева")
			stopAsking = true
		} else {
			fmt.Printf("---> помилка вводу. Оберіть варіант: 1 - будиночок лісника; 2 - озеро; 3 - зачароване дерево.\n")
			fmt.Scan(&whereToGo)
			stopAsking = false
		}
	}

	if whereToGo == 1 {
		fmt.Println("\n\nАліса підійшла до будиночку. Двері масивні та старі і Алісі не вистачило сили їх відкрити. Була лише невелика щілина, через яку було видно охайну кухню з безліччу баночок з травами. Видно було, що в будинку давно ніхто не жив. Через щілину Алісі ніяк не пролізти. І тут вона відчула щось у кармані. Можливо вона знайде те, що допоможе їй зайти в будинок.")

		checkPokets()

		fmt.Printf("\n ---> використайте один з предметів знайдених у карманах щоб зайти до будинку. \nвведіть 1 - щоб використати печиво, яке робить вас меньшим; \n введіть 2 - щоб використати печиво, яке робить вас більшим; \n введіть 3 - щоб використати печиво, яке робить вас зеленим; \n введіть 4 - щоб використати сірники.")
		var useItem int
		fmt.Scan(&useItem)

		var finish bool
		for !finish {
			switch {
			case useItem == pocketItemToBeSmall:
				fmt.Println("Ви використали печиво, яке робить вас меньшим і змогли зайти в будиночок де Аліса знайшла карту, яка допомогла їй вийти з лісу. Вітання.")
				finish = true
			case useItem == pocketItemToBeBig:
				fmt.Println("Ви використали печиво, яке робить вас більшим і це не допомогло зайти в будиночок. Спробуйте ще.")
				fmt.Scan(&useItem)
			case useItem == pocketItemToBeGreen:
				fmt.Println("Ви використали печиво, яке робить вас зеленим і це не допомогло зайти в будиночок. Спробуйте ще.")
				fmt.Scan(&useItem)
			case useItem == pocketItemMatches:
				fmt.Println("Ви використали сірники, це не допомогло зайти в будиночок. Спробуйте ще.")
				fmt.Scan(&useItem)
			}
		}

	}
	if whereToGo == 2 {
		fmt.Println("\n\nАліса підійшла до озера. Там сиділи три жаби. Вона захотіла спитати їх як вийти з лісу, але вони не розмоаляли ні з ким у лісі хто не був зеленого кольору як вони. 'Що ж робити' подумала Аліса. Аж раптом вона відчула щось у кармані. Можливо вона знайде те, що допоможе їй поговорити з жабами.")

		checkPokets()

		fmt.Printf("\n ---> використайте один з предметів знайдених у карманах щоб поговорити з жабами. \nвведіть 1 - щоб використати печиво, яке робить вас меньшим; \n введіть 2 - щоб використати печиво, яке робить вас більшим; \n введіть 3 - щоб використати печиво, яке робить вас зеленим; \n введіть 4 - щоб використати сірники.")
		var useItem int
		fmt.Scan(&useItem)

		var finish bool
		for !finish {
			switch {
			case useItem == pocketItemToBeSmall:
				fmt.Println("Ви використали печиво, яке робить вас меньшим і це не допомогло поговорити з жабами. Спробуйте ще.")
				fmt.Scan(&useItem)
			case useItem == pocketItemToBeBig:
				fmt.Println("Ви використали печиво, яке робить вас більшим і це не допомогло поговорити з жабами. Спробуйте ще.")
				fmt.Scan(&useItem)
			case useItem == pocketItemToBeGreen:
				fmt.Println("Ви використали печиво, яке робить вас зеленим і жаби згодились показати Алісі дорогу як вийти з лісу. Вітання.")
				finish = true
			case useItem == pocketItemMatches:
				fmt.Println("Ви використали сірники, це не допомогло поговорити з жабами. Спробуйте ще.")
				fmt.Scan(&useItem)
			}
		}

	}

	if whereToGo == 3 {
		fmt.Println("\n\nАліса підійшла до зачарованого дерева та почула як хтось плаче. То була маленька фея ельф, яка розповіла, що подув сильний вітер та потушив ліхтар-свічку і тепер її сестра заблукає і не зможе повернутися додому адже орієнтується на світло від ліхтаря. І Аліса відчула щось у кармані. Можливо вона знайде те, що допоможе вирішити проблему феї.")

		checkPokets()

		fmt.Printf("\n ---> використайте один з предметів знайдених у карманах щоб запалити свічку. \nвведіть 1 - щоб використати печиво, яке робить вас меньшим; \n введіть 2 - щоб використати печиво, яке робить вас більшим; \n введіть 3 - щоб використати печиво, яке робить вас зеленим; \n введіть 4 - щоб використати сірники.")
		var useItem int
		fmt.Scan(&useItem)

		var finish bool
		for !finish {
			switch {
			case useItem == pocketItemToBeSmall:
				fmt.Println("Ви використали печиво, яке робить вас меньшим і це не допомогло запалити свічку. Спробуйте ще.")
				finish = true
			case useItem == pocketItemToBeBig:
				fmt.Println("Ви використали печиво, яке робить вас більшим і це не допомогло запалити свічку. Спробуйте ще.")
				fmt.Scan(&useItem)
			case useItem == pocketItemToBeGreen:
				fmt.Println("Ви використали печиво, яке робить вас зеленим і це не допомогло запалити свічку. Спробуйте ще.")
				fmt.Scan(&useItem)
			case useItem == pocketItemMatches:
				fmt.Println("Ви використали сірники, запалили свічку, а фея заспокоївшись розповіла Алісі як орієнтуючись по зіркам вийти з лісу. Вітання.")
				finish = true
			}
		}
	}

}

type whatIsAround struct {
	enchantedHouse string
	path           string
	lake           string
	tree           string
}

func (w whatIsAround) printArea() {
	fmt.Printf("Вона озирнулася навкруги та побачила вдалині: %s, %s, %s, %s. Куди ж піти?", w.enchantedHouse, w.path, w.lake, w.tree)
}

type coockie struct {
	cookieToBeSmall  string
	coockieToBeBig   string
	coockieToBeGreen string
}

type whatIsInPocket struct {
	coockie
	matches string
}

func (b whatIsInPocket) printPocket() {
	fmt.Printf("В кармані вона знайшла:%s, %s, %s, %s.", b.cookieToBeSmall, b.coockieToBeBig, b.coockieToBeGreen, b.matches)
}

func getUp() {

	fmt.Println("Аліса прокинулась у зачарованому лісі.")

	isAround := whatIsAround{
		enchantedHouse: "будиночок лісника",
		path:           "стежка до озера",
		lake:           "озеро",
		tree:           "зачароване дерево",
	}

	isAround.printArea()

}

func checkPokets() {

	leftPocket := whatIsInPocket{
		coockie: coockie{cookieToBeSmall: "печиво, яке робить тебе меньшим", coockieToBeBig: "печиво, яке робить тебе більшим", coockieToBeGreen: "печиво, яке робить тебе зеленим"},
	}
	rightPocket := whatIsInPocket{
		matches: "сірники",
	}

	fmt.Printf("\n ---> Щоб перевірити лівий карман введіть 'left', щоб перевірити лівий карман введіть 'right'. Введіть 'stop' для завершення")
	var checkPocket string
	fmt.Scan(&checkPocket)

	for checkPocket != "stop" {
		if checkPocket == "left" {
			fmt.Println("\nАліса перевірила лівий карман")
			leftPocket.printPocket()
		} else {
			if checkPocket == "right" {
				fmt.Println("\n---> Аліса перевірила правий карман")
				rightPocket.printPocket()
			}

		}
		fmt.Printf("\n\n перевірте інший карман або завершіть перевірку ввівши 'stop'   ")
		fmt.Scan(&checkPocket)
	}

}
