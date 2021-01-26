# TaskManagementApp
Web app to task manager

HTTP API Documentation

1. /projects                                                 - to show all available projects
2. /project/projectId                                        - to show project with all its columns, tasks and comments. Where projectId is integer number of project_id
3. /create/project/projectName/projectDescription            - to create new project. Where projectName is name of the new project, where is projectDescription is decription text of project
4. /update/project/projectId/projectName/projectDescription  - to modify existing project. Where projectId is project_id of existing project, projectName and projectDescription - new name and descripton
5. /delete/project/projectId                                 - to delete project. Where projectId is project_id of existing project
6. /delete/project/all                                       - to delete all projects. It also deletes all columns, tasks and comments

7. /create/column/columnName/columnStatus/columnListPos/projectId           - to create new column. Where columnName is new name, columnStatus is status of column, columnListPos is column's position (it displays in ascending order), projectId is id of project to which column is assigned
8. /update/column/columnId/columnName/columnStatus/columnListPos/projectId  - to modify existing column. Where columnId is id of column to modify, other values similar to create column
9. /delete/column/columnId                                                  - to delete column. Where columnId is id of column to delete

10. /create/task/taskName/taskDescription/taskPriority/columnId             - to create new task
11. /update/task/taskId/taskName/taskDescription/taskPriority/columnId      - to modify existin task
12. /delete/task/taskId                                                     - to delete specified task

13. /create/comment/commentText/taskId                                      - to create new comment
14. /update/comment/commentId/commentText/taskId                            - to modify existing comment
15. /delete/comment/commentId                                               - to delete specified comment
