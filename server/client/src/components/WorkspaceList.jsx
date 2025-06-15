import TriggerRunButton from './TriggerRunButton'

function WorkspaceList({ workspaces }) {
  return (
    <table className="w-full border">
      <thead>
        <tr>
          <th className="border px-4 py-2">Name</th>
          <th className="border px-4 py-2">Actions</th>
        </tr>
      </thead>
      <tbody>
        {workspaces.map(ws => (
          <tr key={ws.id}>
            <td className="border px-4 py-2">{ws.attributes.name}</td>
            <td className="border px-4 py-2">
              <TriggerRunButton id={ws.id} />
            </td>
          </tr>
        ))}
      </tbody>
    </table>
  )
}

export default WorkspaceList
