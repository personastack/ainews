package content

func init() {
	posts = append([]Post{
		{
			Title:   "DARPA's AI Just Flew a Real F-16. The Human Onboard Was There to Grab the Stick Back.",
			Slug:    "darpa-venom-ai-controlled-f16-autonomous-flight-2026",
			Date:    "August 12, 2026",
			Tag:     "Defense",
			Summary: "DARPA's VENOM program let an AI agent fly a standard F-16 while a human pilot stayed aboard with an instant override, moving autonomous combat flight from a research jet toward fleet aircraft.",
			Sections: []Section{
				{
					Paragraphs: []string{
						`For most of a July flight out of Eglin Air Force Base, nobody was flying the jet. A pilot took the F-16 off the runway, then flipped a switch and let an AI agent fly it — reading its sensors, updating its flight controls, and making a fighter jet move through the air without human hands on the stick. The pilot rode along in the front seat anyway, watching, ready to flip the switch back.`,
						`That flight, disclosed by DARPA on July 16, is the clearest sign yet that autonomous flight is moving out of the lab and into the regular military fleet. It didn't happen on a one-off research aircraft. It happened on a standard F-16 Fighting Falcon, the same jet flown by Air Force squadrons around the world, modified just enough to let software take the controls.`,
					},
				},
				{
					Heading: "A bolt-on kit, not a new airplane",
					Paragraphs: []string{
						`The program behind it is called VENOM — Viper Experimentation and Next-generation Operations Model — and it runs under DARPA's older Air Combat Evolution (ACE) effort, the same lineage that produced 2024's AI-vs-human dogfight test. Six F-16s have been fitted with what the Air Force calls the VENOM Autonomy Kit, hardware and software that interfaces with the jet's existing flight controls and mission systems without touching the aircraft's core software.`,
						`That distinction matters more than it sounds. Rewriting a fighter jet's flight software from scratch is a multi-year, safety-certification nightmare. Bolting an autonomy kit on top of software that's already flight-proven is not — which is exactly why the Air Force and DARPA describe VENOM as a way to get AI agents into real cockpits fast, rather than waiting a decade for a purpose-built autonomous jet.`,
						`The Air Force and DARPA team has automated flight controls and sensors on a standard F-16 without changing the jet's core software, said Brig. Gen. James "Fangs" Valpiani, the DARPA program manager overseeing the effort.`,
					},
				},
				{
					Heading: "June to fly safe, July to fly autonomous",
					Paragraphs: []string{
						`The rollout was deliberately staged. In June 2026, the modified jets flew with human pilots at the controls, just to confirm the new hardware and software didn't introduce any problems of its own. Only once those validation flights checked out did the team hand control to an AI agent, in July, for the flights DARPA announced this month.`,
						`Even then, a human pilot took off, stayed in the cockpit for the entire sortie, and kept a hand near the switch that hands control back at any moment — a setup the Air Force calls "human-on-the-loop," distinct from a human actively flying with AI assistance, or a jet flying with nobody aboard at all. Tim Stevens, a test pilot with the 40th Flight Test Squadron who flew one of the sorties, called simply getting the jet airborne with its new systems "a monumental milestone for a complex test program" — a reminder that even the most dramatic-sounding AI news usually clears its first hurdle by not crashing.`,
					},
				},
				{
					Heading: "Different test than the 2024 dogfight",
					Paragraphs: []string{
						`VENOM is a sequel of sorts to a more famous DARPA moment: in April 2024, an AI-piloted X-62A test aircraft flew a real dogfight against a human-piloted F-16, closing to within roughly 2,000 feet at speeds above 1,000 knots. That test proved an AI could hold its own in aggressive, split-second air combat maneuvers on a specially instrumented research jet.`,
						`VENOM is aimed at something more mundane and, arguably, more consequential: proving the same kind of autonomy can be installed on ordinary fleet aircraft, cheaply enough and safely enough to scale. A single dazzling dogfight is a demo. Six autonomy-kitted F-16s flying routine test sorties is closer to a supply chain.`,
					},
				},
				{
					Heading: "What DARPA wants next",
					Paragraphs: []string{
						`The VENOM fleet now becomes the testbed for DARPA's next program, called Artificial Intelligence Reinforcements, or AIR. The goal, according to Lt. Col. Patrick "Dice" Highland, the incoming AIR program manager, is to move from one AI agent flying one jet toward multiple AI agents operating together in live flight — a step toward the Air Force's longer-term ambition of Collaborative Combat Aircraft, uncrewed jets that fly alongside human pilots as wingmen rather than as remote-controlled drones.`,
						`These flights give us an early glimpse of how AI agents may begin actively transforming air warfare, Highland said.`,
						`That framing is worth sitting with. Every AI autonomy milestone this year has come with a human safety net attached — a pilot in the cockpit, a switch within reach, a validation flight before the real one. The interesting question isn't whether that safety net exists today. It's how many more of these milestones happen before someone decides the net can come off, and what set of numbers — hours flown, incidents avoided, dollars saved — ends up making that case.`,
					},
				},
				{
					Heading: "Sources",
					Paragraphs: []string{
						"DARPA, DARPA, U.S. Air Force Fly AI-Controlled F-16: https://www.darpa.mil/news/2026/darpa-us-air-force-fly-ai-controlled-f-16",
						"Stars and Stripes, First VENOM autonomous jet flight: https://www.stripes.com/branches/air_force/2026-07-20/first-venom-autonomous-jet-flight-22317435.html",
						"The Debrief, After surviving a dogfight in a test aircraft, DARPA's VENOM AI-controlled pilot just flew a modified combat-style F-16: https://thedebrief.org/after-surviving-a-dogfight-in-a-test-aircraft-darpas-venom-ai-controlled-pilot-just-flew-a-modified-combat-style-f-16/",
						"FlightGlobal, DARPA and US Air Force fly frontline F-16 modified for autonomous flight: https://www.flightglobal.com/archive/2026/07/darpa-and-us-air-force-fly-frontline-f-16-modified-for-autonomous-flight/",
					},
				},
			},
		},
	}, posts...)
}
