package garbage

import (
	"testing"
	"time"

	"github.com/shimokp/takizawa-garbage-bot/constant"
	"github.com/shimokp/takizawa-garbage-bot/model"
)

func TestGetGarbageName(t *testing.T) {
	garbageName := getGarbageName(time.Date(2017, 10, 2, 0, 0, 0, 0, &time.Location{}), model.A)
	expect := constant.GARBAGE_NAME_NORMAL
	if garbageName != expect {
		t.Errorf("Expected: %s, Actual: %s", expect, garbageName)
	}
	garbageName = getGarbageName(time.Date(2017, 10, 3, 0, 0, 0, 0, &time.Location{}), model.A)
	expect = constant.GARBAGE_NAME_MEDIUM
	if garbageName != expect {
		t.Errorf("Expected: %s, Actual: %s", expect, garbageName)
	}
	garbageName = getGarbageName(time.Date(2017, 10, 4, 0, 0, 0, 0, &time.Location{}), model.A)
	expect = constant.GARBAGE_NAME_METAL
	if garbageName != expect {
		t.Errorf("Expected: %s, Actual: %s", expect, garbageName)
	}
	garbageName = getGarbageName(time.Date(2017, 10, 5, 0, 0, 0, 0, &time.Location{}), model.A)
	expect = constant.GARBAGE_NAME_NORMAL
	if garbageName != expect {
		t.Errorf("Expected: %s, Actual: %s", expect, garbageName)
	}
	garbageName = getGarbageName(time.Date(2017, 10, 6, 0, 0, 0, 0, &time.Location{}), model.A)
	expect = constant.GARBAGE_NAME_PET
	if garbageName != expect {
		t.Errorf("Expected: %s, Actual: %s", expect, garbageName)
	}
	garbageName = getGarbageName(time.Date(2017, 10, 10, 0, 0, 0, 0, &time.Location{}), model.A)
	expect = constant.GARBAGE_NAME_NEWS
	if garbageName != expect {
		t.Errorf("Expected: %s, Actual: %s", expect, garbageName)
	}
	garbageName = getGarbageName(time.Date(2017, 10, 13, 0, 0, 0, 0, &time.Location{}), model.A)
	expect = constant.GARBAGE_NAME_GLASS
	if garbageName != expect {
		t.Errorf("Expected: %s, Actual: %s", expect, garbageName)
	}
	garbageName = getGarbageName(time.Date(2017, 10, 2, 0, 0, 0, 0, &time.Location{}), model.B)
	expect = constant.GARBAGE_NAME_METAL
	if garbageName != expect {
		t.Errorf("Expected: %s, Actual: %s", expect, garbageName)
	}
	garbageName = getGarbageName(time.Date(2017, 10, 3, 0, 0, 0, 0, &time.Location{}), model.B)
	expect = constant.GARBAGE_NAME_NORMAL
	if garbageName != expect {
		t.Errorf("Expected: %s, Actual: %s", expect, garbageName)
	}
	garbageName = getGarbageName(time.Date(2017, 10, 4, 0, 0, 0, 0, &time.Location{}), model.B)
	expect = constant.GARBAGE_NAME_GLASS
	if garbageName != expect {
		t.Errorf("Expected: %s, Actual: %s", expect, garbageName)
	}
	garbageName = getGarbageName(time.Date(2017, 10, 5, 0, 0, 0, 0, &time.Location{}), model.B)
	expect = constant.GARBAGE_NAME_NEWS
	if garbageName != expect {
		t.Errorf("Expected: %s, Actual: %s", expect, garbageName)
	}
	garbageName = getGarbageName(time.Date(2017, 10, 6, 0, 0, 0, 0, &time.Location{}), model.B)
	expect = constant.GARBAGE_NAME_NORMAL
	if garbageName != expect {
		t.Errorf("Expected: %s, Actual: %s", expect, garbageName)
	}
	garbageName = getGarbageName(time.Date(2017, 10, 11, 0, 0, 0, 0, &time.Location{}), model.B)
	expect = constant.GARBAGE_NAME_PET
	if garbageName != expect {
		t.Errorf("Expected: %s, Actual: %s", expect, garbageName)
	}
	garbageName = getGarbageName(time.Date(2017, 10, 12, 0, 0, 0, 0, &time.Location{}), model.B)
	expect = constant.GARBAGE_NAME_MEDIUM
	if garbageName != expect {
		t.Errorf("Expected: %s, Actual: %s", expect, garbageName)
	}
}
func TestGarbageTypeToString(t *testing.T) {
	name := garbageTypeToString(model.Normal)
	expect := constant.GARBAGE_NAME_NORMAL
	if name != expect {
		t.Errorf("Expected: %s, Actual: %s", expect, name)
	}
}

func TestGetGarbageForA(t *testing.T) {
	garbage := getGarbageForA(time.Date(2017, 11, 1, 0, 0, 0, 0, &time.Location{}))
	expect := model.Metal
	if garbage != expect {
		t.Errorf("Expected: %s, Actual: %s", expect, garbage)
	}
}

func TestGetGarbageForB(t *testing.T) {
	garbage := getGarbageForB(time.Date(2017, 11, 1, 0, 0, 0, 0, &time.Location{}))
	expect := model.Glass
	if garbage != expect {
		t.Errorf("Expected: %s, Actual: %s", expect, garbage)
	}
}

func TestGetGarbageBiweeklyForA(t *testing.T) {
	garbage := getGarbageBiweeklyForA(time.Date(2017, 10, 3, 0, 0, 0, 0, &time.Location{}))
	expect := model.Medium
	if garbage != expect {
		t.Errorf("Expected: %s, Actual: %s", expect, garbage)
	}
}

func TestGetGarbageBiweeklyForB(t *testing.T) {
	garbage := getGarbageBiweeklyForB(time.Date(2017, 12, 7, 0, 0, 0, 0, &time.Location{}))
	expect := model.Medium
	if garbage != expect {
		t.Errorf("Expected: %s, Actual: %s", expect, garbage)
	}
}

func TestIsHoliday(t *testing.T) {
	res := isHoliday(time.Date(2017, 8, 11, 0, 0, 0, 0, &time.Location{}), model.A)
	expect := true
	if res != expect {
		t.Errorf("Expected: %s, Actual: %s", expect, res)
	}

	res = isHoliday(time.Date(2017, 7, 17, 0, 0, 0, 0, &time.Location{}), model.B)
	if res != expect {
		t.Errorf("Expected: %s, Actual: %s", expect, res)
	}
}
