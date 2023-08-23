# Калькулятор с поддержкой арабских и римских чисел

## Описание задачи

Создайте консольное приложение "Калькулятор". Приложение должно читать из консоли введенные пользователем строки, числа и арифметические операции между ними, а затем выводить в консоль результат выполнения операции.

## Требования:
1. Калькулятор умеет выполнять операции сложения, вычитания, умножения и деления с двумя числами: a + b, a - b, a * b, a / b. Данные передаются в одну строку (смотри пример ниже). Решения, в которых каждое число и арифметическая операция передаются с новой строки, считаются неверными.
2. Калькулятор умеет работать как с арабскими (1,2,3,4,5…), так и с римскими (I,II,III,IV,V…) числами.
3. Калькулятор должен принимать на вход числа от 1 до 10 включительно, не более. На выходе числа не ограничиваются по величине и могут быть любыми.
4. Калькулятор умеет работать только с целыми числами.
5. Калькулятор умеет работать только с арабскими или римскими цифрами одновременно, при вводе пользователем строки вроде 3 + II калькулятор должен указать на ошибку и прекратить работу.
6. При вводе римских чисел ответ должен быть выведен римскими цифрами, соответственно, при вводе арабских — ответ ожидается арабскими.
7. При вводе пользователем не подходящих чисел приложение выводит ошибку в терминал и завершает работу.
8. При вводе пользователем строки, не соответствующей одной из вышеописанных арифметических операций, приложение выводит ошибку и завершает работу.
9. Результатом операции деления является целое число, остаток отбрасывается.
10. Результатом работы калькулятора с арабскими числами могут быть отрицательные числа и ноль. Результатом работы калькулятора с римскими числами могут быть только положительные числа, если результат работы меньше единицы, программа должна указать на исключение.

## Что было сделано
Этот калькулятор предоставляет возможность выполнения арифметических операций как с арабскими, так и с римскими числами. Программа принимает арифметические выражения от пользователя и выполняет заданное действие.

## Операции

Программа поддерживает следующие арифметические операции:

- **Сложение (+)**: Складывает два числа.
- **Вычитание (-)**: Вычитает второе число из первого.
- **Умножение (*)**: Умножает два числа.
- **Деление (/)**: Делит первое число на второе.

## Использование

1. Запустите программу.
2. Введите арифметическое выражение в формате: `операнд оператор операнд`. Примеры: `5 + 3` или `IV * III`.
3. Нажмите клавишу Enter.
4. Программа выполнит операцию и выведет результат в консоли.

## Ограничения

- Программа работает с числами от 1 до 10 включительно.
- Для римских чисел используйте символы: `I`, `II`, `III`, `IV`, `V`, `VI`, `VII`, `VIII`, `IX` и `X`.

## Примеры

### Пример 1

**Ввод**:
V + III

**Вывод**:
VIII

## Заметки

- При вводе арабских чисел используйте значения от 1 до 10.
- Для римских чисел используйте символы от `I` до `X`.
- Программа не поддерживает дробные числа и отрицательные результаты.

---
