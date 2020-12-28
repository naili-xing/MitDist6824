package mr

import "time"



func (m *Master) mapTaskMonitor(FileName string, workerId int)  {
	/*
		The master can't reliably distinguish between crashed workers,
		workers that are alive but have stalled for some reason, and workers
		that are executing but too slowly to be useful. The best you can do is
		have the master wait for some amount of time, and then give up and re-issue
		the task to a different worker. For this lab, have the master wait for ten seconds;
		after that the master should assume the worker has died (of course, it might not have).
	*/
	for i := 1; i <= 10; i++ {
		time.Sleep(time.Second)
		m.Lock()
		if v, ok := m.MapTaskStatus[FileName]; ok && v==StatusFinish{
			DPrintf("[Worker]: one Map Task Finished\n")
			m.Unlock()
			return
		}
		m.Unlock()
	}

	m.Lock()
	DPrintf("[Worker]: one Map Task Failed, ready to be re-assigned\n")
	m.MapFiles = append(m.MapFiles, FileName)
	m.workers[workerId] = WorkerDelay
	m.Unlock()
	return
}


func (m *Master) isMapFinish() bool{

	// number of task is less than data slices
	if len(m.MapTaskStatus) < m.M{
		return false
	}

	for _, v:= range m.MapTaskStatus{
		if v==StatusFinish{
			continue
		}else{
			return false
		}
	}
	return true

}