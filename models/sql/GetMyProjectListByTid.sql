select * from project where teacher_id = 1
                        and published = true
                        and closed = false
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
                        and  project.project_title like '%神奇%'

order by create_at desc limit 0, 10