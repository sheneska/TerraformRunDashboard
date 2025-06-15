import axios from 'axios'
import { useState } from 'react'

function TriggerRunButton({ id }) {
  const [status, setStatus] = useState(null)

  const handleClick = async () => {
    setStatus('triggering...')
    try {
      await axios.post(`http://localhost:8080/workspaces/${id}/run`)
      setStatus('✅ triggered')
    } catch (err) {
      console.error(err)
      setStatus('❌ error')
    }
  }

  return (
    <div>
      <button onClick={handleClick} className="bg-blue-600 text-white px-3 py-1 rounded">
        Trigger Run
      </button>
      {status && <div className="text-sm mt-1">{status}</div>}
    </div>
  )
}

export default TriggerRunButton
