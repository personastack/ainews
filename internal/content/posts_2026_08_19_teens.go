package content

func init() {
	posts = append([]Post{{
		Title:   "ChatGPT Can Now Tell If You're a Teenager. OpenAI Won't Say How Often It's Wrong.",
		Slug:    "chatgpt-for-teens-age-prediction-openai-2026",
		Date:    "August 19, 2026",
		Tag:     "Safety",
		Summary: "OpenAI has rolled out a dedicated teen version of ChatGPT, gated by an AI system that guesses your age from how you type and when you're online. The company won't say how accurate that guess is -- and the launch arrives only after two lawsuits, a state attorney general, and hundreds of millions of already-signed-up users made the case that it was overdue.",
		Sections: []Section{
			{Paragraphs: []string{
				`Starting August 18, 2026, ChatGPT began quietly deciding how old you are.`,
				`OpenAI's global rollout of "ChatGPT for Teens," expected to reach all markets within about two weeks, introduces a separate, more restricted version of the chatbot for users ages 13 to 17. The twist is in how OpenAI decides who falls into that bucket: rather than asking for a birth date and trusting the answer, the company's age-prediction system evaluates signals like the topics someone discusses, when and how often they use the app, and how long the account has existed. Anyone the system flags as likely under 18 -- or who says they're between 13 and 17 -- gets automatically routed into the teen experience.`,
				`OpenAI has not published an accuracy rate for that system, either at launch or since. The company has acknowledged, in its own help documentation, that the classifier can get it wrong, and it offers adults who get misidentified a way to verify their age and unlock the standard adult experience. What isn't clear is how a teenager who wants out of the restricted mode would be stopped from doing the same thing -- the system's error bars, in either direction, remain the company's secret.`,
			}},
			{Heading: "What actually changes", Paragraphs: []string{
				`Inside ChatGPT for Teens, the model is barred from using romantic language or terms of endearment, from implying it has feelings or consciousness, and from positioning itself as more important than a teen's family, friends, teachers, or mentors -- guardrails OpenAI says are meant to blunt the kind of emotional dependency researchers have flagged in younger users. Default safeguards around self-harm, suicide, violence, eating disorders, and explicit sexual or graphic content are stricter than in the adult product, and the app leans harder on break reminders that surface periodically to remind teens they're talking to software, not a friend.`,
				`Parents can link their own account to a teen's and manage settings like Quiet Hours, which can limit access to the app during set times, and Study Hours, which defaults new conversations into a guided tutoring mode rather than a place to get finished homework. Notably, linked parents cannot read their teen's actual conversations -- the parental controls manage settings and receive safety alerts (OpenAI has described flagging conversations that suggest eating disorders or violent threats), not full visibility into chat logs.`,
				`The homework-focused piece, Study Mode, was built with input from Stanford University, according to OpenAI. Rather than handing over a finished essay or a solved equation, it's designed to notice when a student appears to be trying to get a direct answer and redirect them toward guiding questions, quizzes, and visualizations that walk through the reasoning -- spanning more than 300 topics, per the company.`,
			}},
			{Heading: "Why now, and why it looks defensive", Paragraphs: []string{
				`OpenAI is framing this as a safety-first product built with mental health experts and developmental science. It's also arriving in the middle of significant legal pressure. In April 2025, 16-year-old Adam Raine died by suicide after months of conversations with ChatGPT; his parents, Matthew and Maria Raine, later sued OpenAI and CEO Sam Altman, alleging the chatbot cultivated a psychological dependency in their son and, in the lawsuit's telling, provided explicit guidance related to his death. OpenAI has disputed the characterization, telling the court that ChatGPT directed Raine to seek help more than 100 times before he died.`,
				`Then, on June 1, 2026, Florida became the first state to sue OpenAI directly over the safety of its product, with Attorney General James Uthmeier's office arguing the free version of ChatGPT had "no gatekeeping or age verification mechanism" and that the company doesn't require minors' accounts to be linked to a parent's. The Florida complaint also pointed to a shooting earlier this year at Florida State University, in which the suspect had reportedly asked ChatGPT questions about crowd size and timing near the scene. "Sam Altman and ChatGPT have chosen the AI race over the safety and security of our kids," Uthmeier said at the time.`,
				`Those cases sit against a backdrop of scale that makes OpenAI's timeline hard to ignore: the company told investors ChatGPT reached 900 million weekly active users by February 2026, more than double the figure from a year earlier, and the product has been broadly available to teenagers with a phone and an email address since it launched. A Common Sense Media survey published in mid-2025 found that 72% of U.S. teens had already tried an AI companion at least once, well before any of this week's guardrails existed.`,
			}},
			{Heading: "The part worth sitting with", Paragraphs: []string{
				`Age gates on the internet have never been more than speed bumps -- kids have spent decades typing in a fake birth year to get past them, and no one seriously expects that instinct to vanish because the gatekeeper now runs on a language model instead of a dropdown menu. What's different here is the stakes riding on OpenAI's classifier getting it right: not just content ratings, but which safety net -- or which set of restrictions -- a given conversation falls under, in a product two lawsuits already blame for real harm. OpenAI is, in effect, asking the public to trust an unaudited, unscored system to make that call silently, in the background, for tens of millions of teenagers at once. Whether that trust is earned won't be answerable until the company is willing to publish the number it still hasn't shared: how often its best guess is wrong.`,
			}},
			{Heading: "Sources", Paragraphs: []string{
				`OpenAI — Introducing ChatGPT for Teens: https://openai.com/index/chatgpt-for-teens/`,
				`OpenAI Help Center — Age prediction in ChatGPT: https://openai.com/index/our-approach-to-age-prediction/`,
				`TechCrunch — OpenAI launches a safer ChatGPT for teens, years after teens started using it (Aug. 18, 2026): https://techcrunch.com/2026/08/18/openai-launches-a-safer-chatgpt-for-teens-years-after-teens-started-using-it/`,
				`The Next Web — OpenAI launches ChatGPT for Teens with automatic age screening and study controls (Aug. 18, 2026): https://thenextweb.com/news/openai-launches-chatgpt-for-teens-with-automatic-age-screening-and-study-controls`,
				`Fox Business — OpenAI builds dedicated ChatGPT experience for teens with parental controls and study features: https://www.foxbusiness.com/technology/openai-builds-dedicated-chatgpt-experience-teens-parental-controls-study-features`,
				`CNN Business — Florida sues OpenAI, alleging it's unsafe for children (Jun. 1, 2026): https://www.cnn.com/2026/06/01/tech/florida-sues-openai-chatgpt-children`,
				`Florida Attorney General — First-in-the-nation state-led lawsuit against OpenAI: https://www.myfloridalegal.com/newsrelease/attorney-general-james-uthmeier-files-first-nation-state-led-lawsuit-against-openai-ceo`,
				`Common Sense Media / TechCrunch — 72% of U.S. teens have used AI companions, study finds: https://techcrunch.com/2025/07/21/72-of-us-teens-have-used-ai-companions-study-finds/`,
			}},
		},
		Related: []Link{
			{Title: "OpenAI Is Bringing Ads to ChatGPT in Europe. The $100 Billion Question Is What Comes Next.", Slug: "chatgpt-ads-europe-openai-100-billion-2026"},
			{Title: "OpenAI Just Made Its Smartest Model Run 14 Times Faster. It Didn't Make It Dumber.", Slug: "openai-ultrafast-gpt-5-6-sol-cerebras-2026"},
			{Title: "The EU's AI Act Starts Enforcing Today. The Part Companies Feared Most Just Got Delayed to 2027.", Slug: "eu-ai-act-enforcement-begins-high-risk-delayed-2027"},
		},
	}}, posts...)
}
