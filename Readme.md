This is a fairly simple little Command Line TODO list

I felt like a simple tool that keeps track of things that I need to finish up on a daily basis would be an interesting 
task.

Backend is a Postgres database, which might be overkill, but I just like Postgres.

Single user initially

You can create a task, update a description on a task, mark it as complete, and view all tasks

SQL is attached

Uses PGX because it's a good, fast golang postgres driver, and it doesn't go overboard by adding ORM features
which just aren't necessary for this application

Table structure is:

Task:
id autoincrement
taskname varchar
description string
date_added date
date_updated date
is_complete bool