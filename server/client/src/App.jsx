import { useEffect, useState } from 'react'
import axios from 'axios'
import WorkspaceList from './components/WorkspaceList'

function App() {
  const [workspaces, setWorkspaces] = useState([])

  useEffect(() => {
    axios.get('http://localhost:8080/workspaces')
      .then(res => setWorkspaces(res.data.data))
      .catch(err => console.error(err))
  }, [])

  return (
    <div className="p-6 max-w-4xl mx-auto">
      <h1 className="text-2xl font-bold mb-4">Terraform Run Dashboard</h1>
      <WorkspaceList workspaces={workspaces} />
    </div>
  )
}

export default App
