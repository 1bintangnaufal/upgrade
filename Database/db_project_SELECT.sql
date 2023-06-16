SELECT tb_project.id, author_id, project_title, start_date, finish_date, description, toggle_a, toggle_b, toggle_c, toggle_d, image, tb_user.first_name as author
FROM tb_project 
JOIN tb_user 
ON tb_project.author_id = tb_user.id
ORDER BY tb_user.id;