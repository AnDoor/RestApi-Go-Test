
nota_parcial('Ramon', 1, 90).
nota_parcial('Ramon', 2, 88).
nota_parcial('Raul', 1, 45).
nota_parcial('Raul', 2, 55).
nota_parcial('Priscilla', 1, 85).
nota_parcial('Priscilla', 2, 92).
nota_parcial('Carol', 1, 70).
nota_parcial('Carol', 2, 65).
nota_parcial('Sheyna', 1, 80).
nota_parcial('Sheyna', 2, 78).
nota_parcial('David', 1, 60).    
nota_parcial('David', 2, 75).
nota_parcial('Elena', 1, 95).    
nota_parcial('Elena', 2, 90).
nota_parcial('Fernando', 1, 30).
nota_parcial('Fernando', 2, 40).
nota_parcial('Gabriela', 1, 68). 
nota_parcial('Gabriela', 2, 72).
nota_parcial('Hector', 1, 55).   
nota_parcial('Hector', 2, 65).
nota_parcial('Isabel', 1, 48).  
nota_parcial('Isabel', 2, 70).
nota_parcial('Juan', 1, 82).     
nota_parcial('Juan', 2, 85).
nota_parcial('Karla', 1, 52).    
nota_parcial('Karla', 2, 58).
nota_parcial('Luis', 1, 20).     
nota_parcial('Luis', 2, 30).
nota_parcial('Maria', 1, 75).
nota_parcial('Maria', 2, 70).



% Regla: promedio_estudiante(Nombre, Promedio)
% Calcula el promedio de las dos calificaciones de un estudiante.
promedio_estudiante(Nombre, Promedio) :-
    nota_parcial(Nombre, 1, Nota1),
    nota_parcial(Nombre, 2, Nota2),
    Promedio is (Nota1 + Nota2) / 2.

% Regla: estado_final(Nombre, Estado)
% Determina si un estudiante "Aprobó" o "Reprobó" el curso basándose en el promedio.
estado_final(Nombre, 'Aprobado') :-
    promedio_estudiante(Nombre, Promedio),
    Promedio >= 60.

estado_final(Nombre, 'Reprobado') :-
    promedio_estudiante(Nombre, Promedio),
    Promedio < 60.

% Regla: necesita_recuperacion(Nombre)
% Identifica a los estudiantes que obtuvieron menos de 50 puntos en CUALQUIERA de los parciales.
necesita_recuperacion(Nombre) :-
    nota_parcial(Nombre, 1, Nota1),
    Nota1 < 50.

necesita_recuperacion(Nombre) :-
    nota_parcial(Nombre, 2, Nota2),
    Nota2 < 50.