select * from (
    select * from project
        inner join learn_project on (
            learn_project.student_id = 1 and
            learn_project.learning = true and
            project.id = learn_project.project_id
        )
) as project where true
    and exists (
        select project_subject.project_id from project_subject
            where project_subject.project_id = project.id and (
                project_subject.subject = '英语' or project_subject.subject = '数学'
            )
    )
    and exists (
        select project_skill.project_id from project_skill
            where project_skill.project_id = project.id and (
                project_skill.skill = '生活与职业技能'
            )
    )
    and project.project_title like '%神奇%'

order by create_at desc limit 0, 10