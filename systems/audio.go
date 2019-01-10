package systems

import (
	"github.com/andygeiss/ecs"
	"github.com/andygeiss/ecs/components"
	"github.com/gen2brain/raylib-go/raylib"
)

// Audio ...
type Audio struct {
	sounds map[string]rl.Sound
}

// NewAudio ...
func NewAudio() ecs.System {
	return &Audio{
		sounds: map[string]rl.Sound{},
	}
}

// Process ...
func (s *Audio) Process(entityManager *ecs.EntityManager) {
	for _, e := range entityManager.FilterBy("sound") {
		sound := e.Get("sound").(*components.Sound)
		fileName := sound.Filename
		if fileName == "" {
			continue
		}
		// Playing a sound should be event-based.
		// If a ball has a collision with a player or enemy paddle, then there should be a "collision" sound.
		// Therefore a sound contains a map of event/filename pairs.
		// Preload the sound first, if the file is not currently loaded.
		if _, exists := s.sounds[fileName]; !exists {
			s.sounds[fileName] = rl.LoadSound(fileName)
		}
		// Now Play the sound in the background.
		go func() {
			snd := s.sounds[fileName]
			if !rl.IsSoundPlaying(snd) {
				rl.PlaySound(snd)
			}
			// Prevent infinity sound-loop ;-)
			// The collision system needs to set the Filename to "collision" again.
			sound.Filename = ""
		}()
	}
}

// Setup ...
func (s *Audio) Setup() {
	rl.InitAudioDevice()
}

// Teardown ...
func (s *Audio) Teardown() {
	for _, sound := range s.sounds {
		rl.UnloadSound(sound)
	}
	rl.CloseAudioDevice()
}
