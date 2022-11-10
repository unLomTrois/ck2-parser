# Что это?
Это программа для форматирования языка разработки модов для игры [Crusader Kings 2](https://ck2.paradoxwikis.com/Scripting) 

## Зачем?
Во время разработки [мода](https://github.com/Rystic/ATLA-restored) для этой игры, наша команда столкнулась с тем, что большоие количество людей абсолютно хаотично форматируют код.
У нас уже был [форматтер](https://github.com/cwtools/cwtools-vscode), но нас не устраивали правила его работы, и низкая поддержка проекта.
Мне пришла идея разработать свой форматтер по типу prettier из мира веба.

Однако это оказалось сложнее простого добавления табов, замены пробелов на табы, и проч. 

Поэтому пришлось также разработать парсер языка, благо он оказался не очень сложным

## Как?

Состоит из двух частей: парсер и форматтер

Парсер преобразует текст файла с кодом в конкретное синтаксическое дерево, для этого он использует скользящий токенизатор

Форматтер (линтер) анализирует получившееся синтаксическое дерево и реконструирует по нему текст кода, раставляет \n, \t, скобки, ключи, значения, блоки, и прочее.
