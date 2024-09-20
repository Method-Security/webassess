package url

import (
	"strings"
)

func CreateHTMLAnalysisPrompt(htmlCode string) string {
	promptParts := []string{
		"Task: Analyze the following HTML code and provide a response in JSON format according to the specified schema.",
		"",
		"Instructions:",
		"1. Summarize the HTML code in terms of its functionality and purpose.",
		"2. Analyze the HTML code for potential vulnerabilities.",
		"3. Check for any potential sensitive data exposed in the code.",
		"4. Provide your analysis in the following JSON format:",
		"",
		"{",
		"  \"codeSummary\": \"A brief summary of the HTML code's functionality and purpose\",",
		"  \"potentialVulnerabilities\": true/false,",
		"  \"vulnerabilitiesSummary\": \"A summary of potential vulnerabilities, if any\",",
		"  \"potentialSensitiveData\": true/false,",
		"  \"sensitiveDataSummary\": \"A summary of potential sensitive data exposed, if any\"",
		"}",
		"",
		"Notes:",
		"- The 'codeSummary' field is required and should always be provided.",
		"- The 'potentialVulnerabilities' and 'potentialSensitiveData' fields are required boolean values.",
		"- If 'potentialVulnerabilities' is true, provide a non-null 'vulnerabilitiesSummary'.",
		"- If 'potentialSensitiveData' is true, provide a non-null 'sensitiveDataSummary'.",
		"- If no vulnerabilities or sensitive data are found, set the respective boolean to false and omit the summary fields.",
		"- Only add the requested JSON output. Do not include any additional information.",
		"",
		"Analyze the following HTML code:",
		"```html",
		htmlCode,
		"```",
		"",
		"Provide your analysis in the specified JSON format:",
	}

	return strings.Join(promptParts, "\n")
}

func CreateHTMLSynthesisPrompt(combinedOutput string) string {
	promptParts := []string{
		"Task: Synthesize the following two JSON outputs from an HTML code analysis into a single, comprehensive analysis.",
		"",
		"Instructions:",
		"1. Combine the information from both analyses, resolving any conflicts or differences.",
		"2. Provide a more detailed and comprehensive analysis based on the combined information.",
		"3. Output the result in the same JSON format as the input, with these fields:",
		"   - codeSummary: A comprehensive summary of the HTML code's functionality and purpose",
		"   - potentialVulnerabilities: true if any vulnerabilities were found in either analysis, otherwise false",
		"   - vulnerabilitiesSummary: A detailed summary of all potential vulnerabilities found (omit if none found)",
		"   - potentialSensitiveData: true if any sensitive data was found in either analysis, otherwise false",
		"   - sensitiveDataSummary: A detailed summary of all potential sensitive data found (omit if none found)",
		"",
		"Here are the two analysis outputs to synthesize:",
		combinedOutput,
		"",
		"Provide your synthesized analysis in the specified JSON format:",
	}

	return strings.Join(promptParts, "\n")
}
