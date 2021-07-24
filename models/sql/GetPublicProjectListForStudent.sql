select * from (
                  select * from project where project.published = true
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
                                          and project.project_title like '%%'
              ) as p1 left join learn_project on (
            p1.id = learn_project.project_id and learn_project.student_id = 1
    )
order by p1.create_at desc limit 0, 10