package main

type (
	//Point struct
	Point struct {
		exists bool
		k      int
	}

	//Lock struct
	Lock struct {
		i  int
		j  int
		p1 Point
		p2 Point
	}
)

//FindLockPoint Sudoku
func (s *Sudoku) FindLockPoint(prevL Lock) (l Lock) {
	for i := prevL.i + 1; i < 9; i++ {
		for j := prevL.j + 1; j < 9; j++ {
			if s.CellNum(i, j) == 0 {
				count := 0
				for k := 0; k < 9; k++ {
					if s.variants[i][j][k] == true {
						count++
						if count > 2 {
							l.p1.exists = false
							l.p2.exists = false
							break
						}
						l.i = i
						l.j = j
						if l.p1.exists == false {
							l.p1.k = k
							l.p1.exists = true
						} else {
							l.p2.k = k
							l.p2.exists = true
						}
					}
				}
				if (count == 2) && (l.p2.exists == true) {
					break
				}
			}
		}
	}
	return
}

//ProcLocks Sudoku
func (s *Sudoku) ProcLocks() bool {
	lock := Lock{-1, -1, Point{false, 0}, Point{false, 0}}
	for {
		lock := s.FindLockPoint(lock)
		if lock.p2.exists == true {
			s1 := s.Clone()
			s1.variants[lock.i][lock.j][lock.p1.k] = false
			if s1.Proc() == true && s1.IsCorrect(s1.output) {
				s.Assign(s1)
				return true
			}

			s2 := s.Clone()
			s2.variants[lock.i][lock.j][lock.p2.k] = false
			if s2.Proc() == true && s2.IsCorrect(s2.output) {
				s.Assign(s2)
				return true
			}
		} else {
			break
		}
		lock.p1.exists = false
		lock.p2.exists = false
	}
	return false
}
