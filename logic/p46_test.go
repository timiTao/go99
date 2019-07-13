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
		if and(true, false) {
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
		if !or(true, false) {
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
		if !nand(true, false) {
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
		if nor(true, false) {
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
		if !xor(true, false) {
			t.Error("Expected to be true")
		}
	})
	t.Run("all false", func(t *testing.T) {
		if xor(false, false) {
			t.Error("Expected to be false")
		}
	})
}
