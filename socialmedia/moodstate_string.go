// Code generated by "stringer -type=MoodState"; DO NOT EDIT

package socialmedia

import "fmt"

const _MoodState_name = "MoodStateNeutralMoodStateHappyMoodStateSadMoodStateAngryMoodStateHopefulMoodStateThrilledMoodStateBoredMoodStateShyMoodStateComicalMoodStateOnCloudNine"

var _MoodState_index = [...]uint8{0, 16, 30, 42, 56, 72, 89, 103, 115, 131, 151}

func (i MoodState) String() string {
	if i < 0 || i >= MoodState(len(_MoodState_index)-1) {
		return fmt.Sprintf("MoodState(%d)", i)
	}
	return _MoodState_name[_MoodState_index[i]:_MoodState_index[i+1]]
}
