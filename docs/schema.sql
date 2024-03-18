create table Students
(
    id   bigint        not null primary key,
    name varchar(1024) not null
);

create table Scores
(
    id         bigint not null primary key,
    semester   int    not null,
    student_id bigint null,
    score      int    not null,
    constraint Scores_Students_id_fk
        foreign key (student_id) references Students (id)
);

