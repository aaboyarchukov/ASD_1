# Task9

Разъяснение:

В данной задаче нам необходимо развернуть список в обратную сторону, то есть, чтобы порядок его элементов был от указателя tail до указателя head. Для этого нам необходимо воспользоваться указателями next и prev (реализовывать мы будем с помощью двунаправленного списка), мы менем местами узлы между указателями prev и next, проходясь по каждому узлу (двигаемся от указателя tail до head), когда доходим до head, то последним действием будет просто поменять указатели head и tail местами.

Алгоритм:

1. Пусть tempNode - указатель на конец списка (list.tail)
2. Проходимся циклом по списку до головы (list.head) и меняем местами ссылку на прошлый узел с ссылкой на следующий узел
3. После прохождения циклом мы меняем местами указатели на head и tail.

# Task10

Разъяснение:

Суть данной задачи сводится к тому, чтобы найти в списке циклы, то есть замкнутые лруг на друге узлы (состояние при котором ссылка на следующий узел указывает на прошлый). Алгоритм заключается в том, чтобы циклом проходится по списку до указателя tail и проверять ссылку на следующий узел от того узла, на который указывает tail, если ссылка на следующий узел дает нам любой из прошлых узлов, значит в списке есть циклы.

Алгоритм:

1. Проверяем, что список не пустой, если пустой, то возращаем false
2. Проверяем случай, когда список состоит из 1 узла, если он указывает сам на себя, то он является циклическим
3. Далее ставим указатель tempNode на head и указатель prev на предыдущий от head узел
4. Циклом проходясь до tail мы проверяем, что если prev не пустой и tempNode.next указывает на prev, тогда в списке есть цикл и мы возращаем true
5. Когда цикл заканчивается мы возвращаем false, так как не нашли циклов внутри списка.

# Task11

Разъяснение:

В данной задаче необходимо отсотировать список, для этого можно воспользоваться одной из простых сортировок (всавкой или пузырьком, например).

Алгоритм:

1. Объявляем первый цикл, который берет опорный элемент, учитываем, что им мы проходимся по списку, пока не дойдем до пустого узла.
2. Далее объявляем цикл, который берет побочные элементы для сравнения. В нем мы будем сравнивать значение текущего узла и следующего, если первый больше или равен второму, что меняем значения узлов местами. Важно помнить на этом шаге, что вторым циклом мы идем до последнего узла, то есть до list.tail, чтобы не наткнуться на пустой узел, который может вызвать ошибку ```<nil reference>```.

# Task12

Разъяснение:

Суть задачи заключается в том, чтобы слить в один список два отсортированных списка, и итоговый список должен быть также отсортирован. Прежде всего, в самом начале функции, мы отсортируем два входных списка, затем мы циклом бцдем проходится по обоим спискам, пока один из них не закончится, сранивая элементы, если узел из одного списка оказался меньше, тогда вставляем узел в итоговый список и сдвигаем этот узел. В итоге у нас получается третий отсортированный список слитый из двух других.

Алгоритм:

1. Сортируем l1 и l2 с помощью метода, который написали раннее
2. Формируем итоговый список l3, чтобы добавлять в него элементы
3. Проходимя циклом устанавливая два указателя на списки tempNode1 = l1.head и tempNode2 = l2.head, пока не дойдем до конца хотя бы у одного из циклов
4. Сравниваем значения узлов двух списков, если значения одного узла будет меньше, то мы добавляем именно этот узел в l3, и сдвигаем указатель на тот список, значение узла которого, было меньшим
5. После прохождения этого цикла, мы организуем два других, которые проходятся по одному и другому списку, но будет из двух циклов работать только тот, который не закончился на п. 4.
6. И из того списка, который не закончился, мы будем добавлть оставшиеся элементы, так как они упорядочены

# Task13

В данном задании необходимо реализовать связный список вместе с dummy nodes (узлы-пустышки). Данный тип узлов гарантирует нам, что мы никогда не выйдем за пределы списка и не наткнемся на пустой указатель, а также наши методы сведутся к одним - универсальным, в которых не надо рассматривать по три случая (вначале, в конце и в середине).

Реализовать это проще всегос помощью флажка внутри структуры DummyNode, так как в Go реализация наследования и работа с объектами оставляет желать лучшего.