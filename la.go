import (
	"context"
	"fmt"
	"io"

	texttospeech "cloud.google.com/go/texttospeech/apiv1"
	"cloud.google.com/go/texttospeech/apiv1/texttospeechpb"
)

// report synthesizes the input string of text or ssml.
func report(w io.Writer, textToSpeech string) error {
	// textToSpeech := "Hello there!"
	ctx := context.Background()

	// Creates a client.
	client, err := texttospeech.NewClient(ctx)
	if err != nil {
		return fmt.Errorf("texttospeech.NewClient: %v", err)
	}
	defer client.Close()

	req := &texttospeechpb.SynthesizeSpeechRequest{
		// Set the text input to be synthesized.
		Input: &texttospeechpb.SynthesisInput{
			InputSource: &texttospeechpb.SynthesisInput_Text{Text: textToSpeech},
		},
		// Build the voice request, select the language code ("en-US") and the ssml
		// voice gender ("neutral").
		Voice: &texttospeechpb.VoiceSelectionParams{
			LanguageCode: "en-US",
			SsmlGender:   texttospeechpb.SsmlVoiceGender_NEUTRAL,
		},
		// Select the type of audio file you want returned.
		AudioConfig: &texttospeechpb.AudioConfig{
			AudioEncoding: texttospeechpb.AudioEncoding_MP3,
		},
	}

	resp, err := client.SynthesizeSpeech(ctx, req)
	if err != nil {
		return fmt.Errorf("SynthesizeSpeech: %v", err)
	}

	// The resp's AudioContent is binary.
	fmt.Fprintln(w, resp.AudioContent)
	return nil
}
  
