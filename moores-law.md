## Moore's law, why is dead

Transistors consume power when they switch and increasing transistor density leads to increased **power** consumption, high power leads to high **temperature**.  
Air cooling (fans) can only remove so much heat.

The dynamic power consumption of CMOS circuits is proportional to frequency (P = α * CFV^2).  
Dennard Scaling law said that voltage should scale with transistor size, and it means that power consumption, and temperature should be remain low. But voltage can't go too low for three main reasons:

- Must stay above threshold voltage (transistors have what's called a threshold voltage. Below a certain voltage, they cannot switch on)
- Noise problems occur
- Does not consider leakage power

