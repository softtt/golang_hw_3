package main

import (
	"fmt"
	"sort"
	"strings"
)

var textToTest = `Как видите, он  спускается  по  лестнице  вслед  за  своим
	другом   Кристофером   Робином,   головой   вниз,  пересчитывая
	ступеньки собственным затылком:  бум-бум-бум.  Другого  способа
	сходить  с  лестницы  он  пока  не  знает.  Иногда ему, правда,
		кажется, что можно бы найти какой-то другой способ, если бы  он
	только   мог   на  минутку  перестать  бумкать  и  как  следует
	сосредоточиться. Но увы - сосредоточиться-то ему и некогда.
		Как бы то ни было, вот он уже спустился  и  готов  с  вами
	познакомиться.
	- Винни-Пух. Очень приятно!
		Вас,  вероятно,  удивляет, почему его так странно зовут, а
	если вы знаете английский, то вы удивитесь еще больше.
		Это необыкновенное имя подарил ему Кристофер  Робин.  Надо
	вам  сказать,  что  когда-то Кристофер Робин был знаком с одним
	лебедем на пруду, которого он звал Пухом. Для лебедя  это  было
	очень   подходящее  имя,  потому  что  если  ты  зовешь  лебедя
	громко: "Пу-ух! Пу-ух!"- а он  не  откликается,  то  ты  всегда
	можешь  сделать вид, что ты просто понарошку стрелял; а если ты
	звал его тихо, то все подумают, что ты  просто  подул  себе  на
	нос.  Лебедь  потом  куда-то делся, а имя осталось, и Кристофер
	Робин решил отдать его своему медвежонку, чтобы оно не  пропало
	зря.
		А  Винни - так звали самую лучшую, самую добрую медведицу
	в  зоологическом  саду,  которую  очень-очень  любил  Кристофер
	Робин.  А  она  очень-очень  любила  его. Ее ли назвали Винни в
	честь Пуха, или Пуха назвали в ее честь - теперь уже никто  не
	знает,  даже папа Кристофера Робина. Когда-то он знал, а теперь
	забыл.
		Словом, теперь мишку зовут Винни-Пух, и вы знаете почему.
		Иногда Винни-Пух любит вечерком во что-нибудь поиграть,  а
	иногда,  особенно  когда  папа  дома,  он больше любит тихонько
	посидеть у огня и послушать какую-нибудь интересную сказку.
		В этот вечер...`

func Top10(text string) []string {

	var top []string

	if text == "" {
		return top
	}

	v := strings.Fields(text)
	sort.Strings(v)

	wordsAndFrequency := getWordsFrequencies(v)
	top = getMostFrequentSortedWords(10, wordsAndFrequency)

	return top
}

func getMostFrequentSortedWords(howMuch int, wordsAndFrequency map[string]int) []string {
	var resultSlice []string

	words := make([]string, 0, len(wordsAndFrequency))
	for name := range wordsAndFrequency {
		words = append(words, name)
	}

	sort.Slice(words, func(i, j int) bool {
		return wordsAndFrequency[words[i]] > wordsAndFrequency[words[j]]
	})

	subSliceLen := 0

	for subSliceLen < howMuch {
		subSlice := getSortedSubSlice(words[subSliceLen:], wordsAndFrequency)
		subSliceLen += len(subSlice)

		resultSlice = append(resultSlice, subSlice...)
	}

	return resultSlice
}

func getWordsFrequencies(v []string) map[string]int {
	wordsAndFrequency := make(map[string]int)

	toCompare := v[0]
	counter := 1
RowLoop:
	for i, s := range v {
		if toCompare == v[i] {
			counter ++
			continue RowLoop
		} else {
			wordsAndFrequency[toCompare] = counter
			toCompare = s
			counter = 1
		}
	}

	return wordsAndFrequency
}

func getSortedSubSlice(words []string, wordsAndFrequency map[string]int) []string {
	var subSlice []string

	for i, name := range words {
		if wordsAndFrequency[words[i]] == wordsAndFrequency[words[i+1]] {
			subSlice = append(subSlice, name)
		} else {
			subSlice = append(subSlice, name)
			words = words[i+1:]
			sort.Strings(subSlice)
			return subSlice
		}
	}
	sort.Strings(subSlice)
	return subSlice
}

func main()  {
	result := Top10(textToTest)

	fmt.Println(result)
}
