package logic

import "testing"

func TestAnd(t *testing.T) {
	t.Run("all true", func(t *testing.T) {
		if !and(true, true) {
			t.Error("Expected to be true")
		}
	})
	t.Run("true/false", func(t *testing.T) {
		if and(true, false) {
			t.Error("Expected to be false")
		}
	})
	t.Run("false/true", func(t *testing.T) {
		if and(false, true) {
			t.Error("Expected to be false")
		}
	})
	t.Run("all false", func(t *testing.T) {
		if and(false, false) {
			t.Error("Expected to be false")
		}
	})
}

func TestOr(t *testing.T) {
	t.Run("all true", func(t *testing.T) {
		if !or(true, true) {
			t.Error("Expected to be true")
		}
	})
	t.Run("true/false", func(t *testing.T) {
		if !or(true, false) {
			t.Error("Expected to be true")
		}
	})
	t.Run("false/true", func(t *testing.T) {
		if !or(false, true) {
			t.Error("Expected to be true")
		}
	})
	t.Run("all false", func(t *testing.T) {
		if or(false, false) {
			t.Error("Expected to be false")
		}
	})
}

func TestNand(t *testing.T) {
	t.Run("all true", func(t *testing.T) {
		if nand(true, true) {
			t.Error("Expected to be false")
		}
	})
	t.Run("true/false", func(t *testing.T) {
		if !nand(true, false) {
			t.Error("Expected to be true")
		}
	})
	t.Run("false/true", func(t *testing.T) {
		if !nand(false, true) {
			t.Error("Expected to be true")
		}
	})
	t.Run("all false", func(t *testing.T) {
		if !nand(false, false) {
			t.Error("Expected to be true")
		}
	})
}

func TestNor(t *testing.T) {
	t.Run("all true", func(t *testing.T) {
		if nor(true, true) {
			t.Error("Expected to be false")
		}
	})
	t.Run("true/false", func(t *testing.T) {
		if nor(true, false) {
			t.Error("Expected to be false")
		}
	})
	t.Run("false/true", func(t *testing.T) {
		if nor(false, true) {
			t.Error("Expected to be false")
		}
	})
	t.Run("all false", func(t *testing.T) {
		if !nor(false, false) {
			t.Error("Expected to be true")
		}
	})
}

func TestXor(t *testing.T) {
	t.Run("all true", func(t *testing.T) {
		if xor(true, true) {
			t.Error("Expected to be false")
		}
	})
	t.Run("true/false", func(t *testing.T) {
		if !xor(true, false) {
			t.Error("Expected to be true")
		}
	})
	t.Run("false/true", func(t *testing.T) {
		if !xor(false, true) {
			t.Error("Expected to be true")
		}
	})
	t.Run("all false", func(t *testing.T) {
		if xor(false, false) {
			t.Error("Expected to be false")
		}
	})
}

func TestImpl(t *testing.T) {
	t.Run("all true", func(t *testing.T) {
		if !impl(true, true) {
			t.Error("Expected to be true")
		}
	})
	t.Run("true/false", func(t *testing.T) {
		if impl(true, false) {
			t.Error("Expected to be false")
		}
	})
	t.Run("false/true", func(t *testing.T) {
		if !impl(false, true) {
			t.Error("Expected to be true")
		}
	})
	t.Run("all false", func(t *testing.T) {
		if !impl(false, false) {
			t.Error("Expected to be true")
		}
	})
}

func TestEqu(t *testing.T) {
	t.Run("all true", func(t *testing.T) {
		if !equ(true, true) {
			t.Error("Expected to be true")
		}
	})
	t.Run("true/false", func(t *testing.T) {
		if equ(true, false) {
			t.Error("Expected to be false")
		}
	})
	t.Run("false/true", func(t *testing.T) {
		if equ(false, true) {
			t.Error("Expected to be false")
		}
	})
	t.Run("all false", func(t *testing.T) {
		if !equ(false, false) {
			t.Error("Expected to be true")
		}
	})
}

func TestPrintLogicTable(t *testing.T) {
	t.Run("print: and", func(t *testing.T) {
		expectedResult := [4][3]bool{  
		   {true, true, true},
		   {true, false, false},
		   {false, true, false},
		   {false, false, false},  
		}
		result := printLogicTable(and)
		if result != expectedResult {
			t.Errorf("Expected to be same array. Expected %v, current: %v", expectedResult, result)
		}
	})
	t.Run("print: impl", func(t *testing.T) {
		expectedResult := [4][3]bool{  
		   {true, true, true},
		   {true, false, false},
		   {false, true, true},
		   {false, false, true},  
		}
		result := printLogicTable(impl)
		if result != expectedResult {
			t.Errorf("Expected to be same array. Expected %v, current: %v", expectedResult, result)
		}
	})
}


