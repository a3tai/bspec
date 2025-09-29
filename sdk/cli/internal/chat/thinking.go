package chat

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/alperdrsnn/clime"
)

// ThinkingAnimationType represents different types of operations
type ThinkingAnimationType int

const (
	ThinkingGeneral ThinkingAnimationType = iota
	ThinkingLLM
	ThinkingFileOps
	ThinkingValidation
	ThinkingSDK
	ThinkingTools
)

// ThinkingAnimation manages an animated thinking indicator
type ThinkingAnimation struct {
	mu          sync.Mutex
	isRunning   bool
	cancel      context.CancelFunc
	animType    ThinkingAnimationType
	customMsg   string
	stepCount   int
	totalSteps  int
}

// NewThinkingAnimation creates a new thinking animation
func NewThinkingAnimation() *ThinkingAnimation {
	return &ThinkingAnimation{}
}

// Start begins the thinking animation
func (t *ThinkingAnimation) Start(animType ThinkingAnimationType, customMsg string) {
	t.mu.Lock()
	defer t.mu.Unlock()

	// Stop any existing animation
	if t.isRunning {
		t.stopUnsafe()
	}

	// Set up new animation
	t.animType = animType
	t.customMsg = customMsg
	t.isRunning = true

	// Create context for cancellation
	ctx, cancel := context.WithCancel(context.Background())
	t.cancel = cancel

	// Start animation goroutine
	go t.animate(ctx)
}

// StartWithProgress begins the thinking animation with progress tracking
func (t *ThinkingAnimation) StartWithProgress(animType ThinkingAnimationType, customMsg string, totalSteps int) {
	t.mu.Lock()
	defer t.mu.Unlock()

	// Stop any existing animation
	if t.isRunning {
		t.stopUnsafe()
	}

	// Set up new animation with progress
	t.animType = animType
	t.customMsg = customMsg
	t.totalSteps = totalSteps
	t.stepCount = 0
	t.isRunning = true

	// Create context for cancellation
	ctx, cancel := context.WithCancel(context.Background())
	t.cancel = cancel

	// Start animation goroutine
	go t.animate(ctx)
}

// UpdateProgress updates the progress counter
func (t *ThinkingAnimation) UpdateProgress(step int) {
	t.mu.Lock()
	defer t.mu.Unlock()

	if t.isRunning {
		t.stepCount = step
	}
}

// IncrementProgress increments the progress counter
func (t *ThinkingAnimation) IncrementProgress() {
	t.mu.Lock()
	defer t.mu.Unlock()

	if t.isRunning {
		t.stepCount++
	}
}

// Stop stops the thinking animation
func (t *ThinkingAnimation) Stop() {
	t.mu.Lock()
	defer t.mu.Unlock()
	t.stopUnsafe()
}

// stopUnsafe stops the animation without locking (internal use only)
func (t *ThinkingAnimation) stopUnsafe() {
	if t.isRunning && t.cancel != nil {
		t.cancel()
		t.isRunning = false

		// Clear the line
		fmt.Print("\r" + string(make([]byte, 80)) + "\r")
	}
}

// animate runs the animation loop
func (t *ThinkingAnimation) animate(ctx context.Context) {
	// Spinner characters for smooth animation
	spinnerChars := []string{"⠋", "⠙", "⠹", "⠸", "⠼", "⠴", "⠦", "⠧", "⠇", "⠏"}
	ticker := time.NewTicker(120 * time.Millisecond)
	defer ticker.Stop()

	frameIndex := 0

	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			t.mu.Lock()

			// Get current spinner character
			spinner := spinnerChars[frameIndex%len(spinnerChars)]
			frameIndex++

			// Build message
			message := t.buildMessage(spinner)

			// Display the animated message
			fmt.Printf("\r%s", clime.Warning.Sprint(message))

			t.mu.Unlock()
		}
	}
}

// buildMessage constructs the appropriate message based on animation type
func (t *ThinkingAnimation) buildMessage(spinner string) string {
	var baseMsg string

	// Use custom message if provided
	if t.customMsg != "" {
		baseMsg = t.customMsg
	} else {
		// Default messages based on type
		switch t.animType {
		case ThinkingLLM:
			baseMsg = "Thinking"
		case ThinkingFileOps:
			baseMsg = "Writing files"
		case ThinkingValidation:
			baseMsg = "Validating documents"
		case ThinkingSDK:
			baseMsg = "Processing with SDK"
		case ThinkingTools:
			baseMsg = "Executing tools"
		default:
			baseMsg = "Working"
		}
	}

	// Add progress if applicable
	if t.totalSteps > 0 {
		return fmt.Sprintf("%s %s (%d/%d)...", spinner, baseMsg, t.stepCount, t.totalSteps)
	}

	return fmt.Sprintf("%s %s...", spinner, baseMsg)
}

// GetMessage returns the current message without spinner (for testing)
func (t *ThinkingAnimation) GetMessage() string {
	t.mu.Lock()
	defer t.mu.Unlock()
	return t.buildMessage("")
}

// IsRunning returns whether the animation is currently running
func (t *ThinkingAnimation) IsRunning() bool {
	t.mu.Lock()
	defer t.mu.Unlock()
	return t.isRunning
}