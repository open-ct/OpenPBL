import React, {useState} from "react";

function LatestProject(obj) {
  const [projects, setProjects] = useState([])


  return (
    <div>
      {projects.map((item, index)=>(
        <div>
        </div>
      ))}
    </div>
  )
}

export default LatestProject