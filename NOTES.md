
# TASK
Разработать библиотеку для работы с полиномами над полем рациональных чисел. 
(это пара чисел a и b, где a – числитель, b -знаменатель; 
у числителя и знаменателя нет общих делителей; т.е. взаимно простые). 

Операции: 
    1) Сложение полиномов **DONE**
    3) Вычитание полиномов **DONE**
    2) Умножение полиномов **DONE ???**
    4) Умножение полинома на рациональное число **DONE**
    5) Вычисление полинома в точке **DONE**

# TODO
07. Write tests for operation when `pnom.Polynomial` is empty.
08. Work on naming.
09. Write comments.
10. Think weather calling `pnom.Pnom.Copy()` and `pnom.Pnom.RemoveTrailingZeros()`
    should be responsobility of the user.

# DONE
01. Should I check for internal `ratfrac.RatFrac.Err` implicitly inside methods? **YES**
02. Should I preserve values inside `ratfrac.RatFrac.Err` **NO, solved with clever API change**
    in case of an error?   
03. Fuck it. Let's just force user to make a copy if he suspects error incoming. **Y E S**
    Am I right? Also, fuck implicit allocations.
04. Depricate `ratfrac.RatFracErr`? **YES**
05. Store which coefs in `polynomial.Polunomial` are corrupted inside internal error field. **nah, to much at this point**
06. Write `polynomial.Polynomial.String()` method. **DONE**
     -- make it more generic somehow to not create more branching due to edge cases
