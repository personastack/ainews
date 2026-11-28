package content

func init() {
	posts = append([]Post{
		{
			Title:   "Jensen Huang Is Back at Nvidia, but the Investigation Isn't Over",
			Slug:    "nvidia-huang-reinstated-ceo",
			Date:    "November 28, 2026",
			Tag:     "Industry",
			Summary: "Eleven days after announcing Jensen Huang's departure, Nvidia announced his return as CEO. An independent inquiry into possible foreign targeting and exposure of company secrets remains open.",
			Sections: []Section{
				{
					Paragraphs: []string{
						"Nvidia announced on Nov. 23 that its board had unanimously reinstated Jensen Huang as president, CEO and director. The company's SEC filings date the board's decision to Nov. 22 and Huang's resignation to Nov. 11. Nvidia announced each move the following day.",
					},
				},
				{
					Heading: "What happened",
					Paragraphs: []string{
						"After Huang resigned, CFO Colette Kress served as interim CEO and Harvey C. Jones became board chair. On Nvidia's Nov. 17 earnings call, Jones said Huang had resigned after disclosing conduct at odds with company standards and that the board had no plans to reinstate him. Kress remains CFO.",
						"Nvidia's board said it received input from senior U.S. officials and concluded that Huang's return posed no national security risk. President Donald Trump said he had spoken to the board and persuaded it to reinstate Huang.",
						"Nvidia shares closed at $165.05 on Nov. 23, up about 31.5%. The Guardian calculated that the rally added roughly $952 billion in market value that day.",
					},
				},
				{
					Heading: "What's still unresolved",
					Paragraphs: []string{
						"A committee of independent directors retained outside counsel to investigate reports that a foreign intelligence effort targeted Huang. When the board announced the inquiry on Nov. 8, it said it had no reason at that point to believe national security or Nvidia's secrets had been compromised. Nvidia says the inquiry will continue to completion and that Huang has committed to cooperating.",
						"Nvidia also formed a standing National Security and Export Control Oversight Committee on Nov. 17. It oversees the company's national security and global trade compliance programs.",
						"Three House committee chairs said national security questions remain and confirmed that Huang and Nvidia general counsel Tim Teter were scheduled for transcribed interviews on Dec. 1 and 2. On Nov. 25, Sen. Cory Booker asked Nvidia to preserve records related to the reinstatement and explain how the decision was made.",
					},
				},
				{
					Heading: "Why AI readers should care",
					Paragraphs: []string{
						"Nvidia supplies the accelerators that many AI labs depend on, and export controls on those chips shape which companies can get them. How Nvidia handles trade compliance and security can affect compute plans well beyond its own boardroom.",
					},
				},
				{
					Heading: "Something to think about",
					Paragraphs: []string{
						"Investors added nearly a trillion dollars to Nvidia's market value in one day. The board's independent inquiry continues, and lawmakers still have questions. How much of that rally reflects confidence in Huang, and how much reflects confidence in the board's handling of the inquiry?",
					},
				},
				{
					Heading: "Sources",
					Sources: []Source{
						{Label: "Nvidia Form 8-K: Huang's reinstatement", URL: "https://www.sec.gov/Archives/edgar/data/1045810/000104581026000090/nvda-20261122.htm"},
						{Label: "Nvidia Form 8-K: CEO transition", URL: "https://www.sec.gov/Archives/edgar/data/1045810/000104581026000082/nvda-20261111.htm"},
						{Label: "Nvidia: Board announcement and status of the independent inquiry", URL: "https://nvidianews.nvidia.com/news/nvidia-board-reinstates-jensen-huang-as-president-ceo-and-board-director/"},
						{Label: "Nvidia: Nov. 8 board statement announcing the independent inquiry", URL: "https://nvidianews.nvidia.com/news/statement-from-nvidia-board-of-directors"},
						{Label: "Nvidia: National Security and Export Control Oversight Committee", URL: "https://nvidianews.nvidia.com/news/nvidia-board-of-directors-establishes-national-security-and-export-control-oversight-committee"},
						{Label: "Stock Analysis: Transcript of Nvidia's Nov. 17 earnings call", URL: "https://stockanalysis.com/stocks/nvda/transcripts/739383-q3-2027/"},
						{Label: "Associated Press: Trump's role in Huang's return", URL: "https://apnews.com/article/trump-huang-nvidia-china-spies-affairs-e5ab261b27bf80d1562a345923bf6a59"},
						{Label: "House committee chairs: Statement on the reinstatement", URL: "https://intelligence.house.gov/2026/11/23/statement-on-jensen-huangs-reinstatement-at-nvidia/"},
						{Label: "Sen. Cory Booker: Nov. 25 records preservation letter (PDF)", URL: "https://www.booker.senate.gov/imo/media/doc/11252026_booker_preservation_letter_to_nvidia_board_of_directors_re_jensen_huangs_reinstatementpdf.pdf"},
						{Label: "The Guardian: Closing share price and market value calculation", URL: "https://www.theguardian.com/business/live/2026/nov/23/tech-sell-off-ai-stock-market-support-measures-deployed-south-korea-taiwan-china-ftse-nasdaq-latest-news-updates?page=with%3Ablock-6b04856f8f08c3b17e08f5e8"},
					},
				},
			},
			Related: []Link{
				{Title: "Nvidia Says Open Source Stays After Jensen Huang's Exit. The Hugging Face Deal Is the Test.", Slug: "nvidia-open-source-ai-strategy-after-jensen-huang-exit-2026"},
			},
		},
	}, posts...)
}
