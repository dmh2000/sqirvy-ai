document.addEventListener('DOMContentLoaded', () => {
    const promptTextarea = document.getElementById('prompt');
    const submitButton = document.getElementById('submit');
    const anthropicResponse = document.getElementById('anthropic-response');
    const openaiResponse = document.getElementById('openai-response');
    const geminiResponse = document.getElementById('gemini-response');

    submitButton.addEventListener('click', async () => {
        const prompt = promptTextarea.value.trim();
        
        if (!prompt) {
            alert('Please enter a prompt');
            return;
        }

        submitButton.disabled = true;
        anthropicResponse.textContent = 'Loading...';
        openaiResponse.textContent = 'Loading...';
        geminiResponse.textContent = 'Loading...';

        try {
            const response = await fetch(`/api/query?prompt=${encodeURIComponent(prompt)}`);
            const data = await response.json();
            
            // Update Anthropic response
            if (data.anthropic.error) {
                anthropicResponse.textContent = `Error: ${data.anthropic.error}`;
            } else {
                anthropicResponse.textContent = data.anthropic.result;
            }

            // Update OpenAI response
            if (data.openai.error) {
                openaiResponse.textContent = `Error: ${data.openai.error}`;
            } else {
                openaiResponse.textContent = data.openai.result;
            }

            // Update Gemini response
            if (data.gemini.error) {
                geminiResponse.textContent = `Error: ${data.gemini.error}`;
            } else {
                geminiResponse.textContent = data.gemini.result;
            }
        } catch (error) {
            anthropicResponse.textContent = `Error: ${error.message}`;
            openaiResponse.textContent = `Error: ${error.message}`;
            geminiResponse.textContent = `Error: ${error.message}`;
        } finally {
            submitButton.disabled = false;
        }
    });
});
