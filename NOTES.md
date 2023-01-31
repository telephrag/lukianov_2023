
# TASK
Разработать библиотеку для работы с полиномами над полем рациональных чисел. 
(это пара чисел a и b, где a – числитель, b -знаменатель; 
у числителя и знаменателя нет общих делителей; т.е. взаимно простые). 

Операции: 
    1) Сложение полиномов
    2) Умножение полиномов
    3) Вычитание полиномов
    4) Умножение полинома на рациональное число
    5) Вычисление полинома в точке

# TODO
05. Store which coefs in `polynomial.Polunomial` are corrupted inside internal error field
06. Write `polynomial.Polynomial.String()` method. 

# DONE
01. Should I check for internal `ratfrac.RatFrac.Err` implicitly inside methods? **YES**
02. Should I preserve values inside `ratfrac.RatFrac.Err` **NO, solved with clever API change**
    in case of an error?   
03. Fuck it. Let's just force user to make a copy if he suspects error incoming. **Y E S**
    Am I right? Also, fuck implicit allocations.
04. Depricate `ratfrac.RatFracErr`? **YES**
