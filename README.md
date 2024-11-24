Я зробив два алгоритми. І ще по три варіанти ціх алгоритмів: звичайна версія, оптимизована певним чином і concurrent версія.
Усі варіанти розклав по пакетам: 
 - optimalblock1 - перший алгоритм
 - optimalblock1opt - перший алгоритм певним чином оптимізований
 - optimalblock1concurrent - concurrent версія
 - optimalblock2 - другий алгоритм
 - optimalblock2opt - другий алгоритм певним чином оптимізований
 - optimalblock2concurrent - concurrent версія

Усі пакети містять функцію `GetOptimalBlock([]map[string]bool) int`, яка повертає шуканий квартал.

Перший алгоритм проходить по масиву кварталів зберігаючи відстані до найближчих інфрастуктурних обʼєктів і максимальну з них відстань. 
Потім знаходить квартал з мінімальною відстанню. Відстань до найближчих об'єктів знаходиться таким чином: пускається "хвиля" вперед 
і назад навколо даного кварталу. І коли хвиля доходить до найближчої інфраструктури, зберігається скільки кварталів вона пройшла.

Другий алгоритм по кожній інфраструктурі відмічає спочатку ті квартали, в яких обʼєкти інфраструктури знаходяться в нульовій віддаленості,
тобто в самому кварталі. Потім навколо кварталів з нулями сусіднім кварталам по одному з кожного боку проставляється відстань 1. Потім 
навколо одиничок проставляються 2 і т.д.

В оптимізованій версії кожного алгоритму додана перевірка на такий квартал, який містить обʼєкти всіх інфраструктур. Якщо такий квартал 
знаходиться, то подальша робота завершується і повертається індекс цього кварталу.

Concurrent-версія першого алгоритму – віддаленість обʼєктів від кожного кварталу вираховується горутинами. Тобто кількість горутин дорівнює 
кількості кварталів. У concurrent-версії другого алгоритму в горутинах відбувається розрахунок відстаней кожного виду інфраструктур. 
Тобто скільки інфраструктур, стільки і горутин.

Написав бенчмарк-тести для кожного варіанту алгоритмів з такими видами вхідних даних: маленький масив кварталів, довгий масив кварталів і 
довгий масив кварталів з "нульовим" кварталом, який містить всі інфраструктури.