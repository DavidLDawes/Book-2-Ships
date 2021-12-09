package main

import (
	"fmt"
	"strconv"

	"fyne.io/fyne"
	"fyne.io/fyne/widget"
)

type weaponsDetail struct {
	missile      int
	beam         int
	pulse        int
	plasma       int
	fusion       int
	sandcaster   int
	accelerator  int
	detailLabels []*widget.Label
	selects      []*widget.Select
	panel        panel
}

var weaponSettings *widget.Form

var (
	detailMissile  *widget.Label = widget.NewLabel("")
	detailBeam     *widget.Label = widget.NewLabel("")
	detailPulse    *widget.Label = widget.NewLabel("")
	detailFusion   *widget.Label = widget.NewLabel("")
	detailSand     *widget.Label = widget.NewLabel("")
	detailPlasma   *widget.Label = widget.NewLabel("")
	detailParticle *widget.Label = widget.NewLabel("")
)

var weaponDetails *widget.Box = widget.NewVBox()

var (
	missileSelect  *widget.Select
	beamSelect     *widget.Select
	pulseSelect    *widget.Select
	fusionSelect   *widget.Select
	sandSelect     *widget.Select
	plasmaSelect   *widget.Select
	particleSelect *widget.Select
)

var weapons weaponsDetail = weaponsDetail{
	missile: 0, beam: 0, pulse: 0, plasma: 0, fusion: 0,
	sandcaster: 0, accelerator: 0,
	detailLabels: []*widget.Label{
		widget.NewLabel("Missile Turrets"),
		widget.NewLabel("Beam Laser Turrets"),
		widget.NewLabel("Pulse Laser Turrets"),
		widget.NewLabel("Fusion Guns"),
		widget.NewLabel("Sandcaster Turrets"),
		widget.NewLabel("Plasma Guns"),
		widget.NewLabel("Particle Beams"),
	},
	selects: []*widget.Select{
		missileSelect, beamSelect, pulseSelect, fusionSelect,
		sandSelect, plasmaSelect, particleSelect,
	},
	panel: panel{},
}

func (w *weaponsDetail) weaponsInit() {
	missileSelect = widget.NewSelect(weaponLevel, nothing)
	missileSelect.SetSelected("0")
	missileSelect.OnChanged = w.missileChanged

	beamSelect = widget.NewSelect(weaponLevel, nothing)
	beamSelect.SetSelected("0")
	beamSelect.OnChanged = w.beamChanged

	pulseSelect = widget.NewSelect(weaponLevel, nothing)
	pulseSelect.SetSelected("0")
	pulseSelect.OnChanged = w.pulseChanged

	fusionSelect = widget.NewSelect(weaponLevel, nothing)
	fusionSelect.SetSelected("0")
	fusionSelect.OnChanged = w.fusionChanged

	plasmaSelect = widget.NewSelect(weaponLevel, nothing)
	plasmaSelect.SetSelected("0")
	plasmaSelect.OnChanged = w.plasmaChanged

	sandSelect = widget.NewSelect(weaponLevel, nothing)
	sandSelect.SetSelected("0")
	sandSelect.OnChanged = w.sandChanged

	particleSelect = widget.NewSelect(weaponLevel, nothing)
	particleSelect.SetSelected("0")
	particleSelect.OnChanged = w.particleChanged

	weaponSettings = widget.NewForm(
		widget.NewFormItem("Missile", missileSelect),
		widget.NewFormItem("Beam", beamSelect),
		widget.NewFormItem("Pulse", pulseSelect),
		widget.NewFormItem("Fusion", fusionSelect),
		widget.NewFormItem("Sand", sandSelect),
		widget.NewFormItem("Plasma", plasmaSelect),
		widget.NewFormItem("Accelerators", particleSelect),
	)
}

func (w *weaponsDetail) weaponsSelectInit() {
	weaponDetails.Children = make([]fyne.CanvasObject, 0)
	// Check each and add if needed
	addWeapon(weapons.missile, weaponDetails, detailMissile)
	addWeapon(weapons.beam, weaponDetails, detailBeam)
	addWeapon(weapons.pulse, weaponDetails, detailPulse)
	addWeapon(weapons.plasma, weaponDetails, detailPlasma)
	addWeapon(weapons.fusion, weaponDetails, detailFusion)
	addWeapon(weapons.sandcaster, weaponDetails, detailSand)
	addWeapon(weapons.accelerator, weaponDetails, detailParticle)
	weaponDetails.Refresh()
}

func (w *weaponsDetail) missileChanged(value string) {
	missiles, err := strconv.Atoi(value)
	if err == nil {
		weapons.missile = missiles
		if w.countWeapons() > hull.maxHP {
			weapons.missile = missiles - w.countWeapons() + hull.maxHP
			if weapons.missile < 0 {
				weapons.missile = 0
			}
			if weapons.missile != missiles {
				missileSelect.SetSelected(strconv.Itoa(weapons.missile))
			}
		}
		w.setWeaponDetails()
	}
	w.buildMissile()
	berths.buildCrew()
	//.buildTotal()
}

func (w *weaponsDetail) beamChanged(value string) {
	beamTurrets, err := strconv.Atoi(value)
	if err == nil {
		weapons.beam = beamTurrets
		if w.countWeapons() > hull.maxHP {
			weapons.beam = beamTurrets - w.countWeapons() + hull.maxHP
			if weapons.beam < 0 {
				weapons.beam = 0
			}
			beamSelect.SetSelected(strconv.Itoa(weapons.beam))
		}
		w.setWeaponDetails()
	}
	w.buildBeam()
	berths.buildCrew()
	//.buildTotal()
}

func (w *weaponsDetail) pulseChanged(value string) {
	pulse, err := strconv.Atoi(value)
	if err == nil {
		weapons.pulse = pulse
		if w.countWeapons() > hull.maxHP {
			weapons.pulse = pulse - w.countWeapons() + hull.maxHP
			if weapons.pulse < 0 {
				weapons.pulse = 0
			}
			pulseSelect.SetSelected(strconv.Itoa(weapons.pulse))
		}
		w.setWeaponDetails()
	}
	w.buildPulse()
	berths.buildCrew()
	//.buildTotal()
}

func (w *weaponsDetail) fusionChanged(value string) {
	fusion, err := strconv.Atoi(value)
	if err == nil {
		weapons.fusion = fusion
		if w.countWeapons() > hull.maxHP {
			weapons.fusion = fusion - w.countWeapons() + hull.maxHP
			if weapons.fusion < 0 {
				weapons.fusion = 0
			}
			fusionSelect.SetSelected(strconv.Itoa(weapons.fusion))
		}
		w.setWeaponDetails()
	}
	w.buildFusion()
	berths.buildCrew()
	//.buildTotal()
}

func (w *weaponsDetail) sandChanged(value string) {
	sand, err := strconv.Atoi(value)
	if err == nil {
		weapons.sandcaster = sand
		if w.countWeapons() > hull.maxHP {
			weapons.sandcaster = sand - w.countWeapons() + hull.maxHP
			if weapons.sandcaster < 0 {
				weapons.sandcaster = 0
			}
			sandSelect.SetSelected(strconv.Itoa(weapons.sandcaster))
		}
		w.setWeaponDetails()
	}
	w.buildSand()
	berths.buildCrew()
	// w.buildTotal()
}

func (w *weaponsDetail) plasmaChanged(value string) {
	plasma, err := strconv.Atoi(value)
	if err == nil {
		weapons.plasma = plasma
		if w.countWeapons() > hull.maxHP {
			weapons.plasma = plasma - w.countWeapons() + hull.maxHP
			if weapons.plasma < 0 {
				weapons.plasma = 0
			}
			plasmaSelect.SetSelected(strconv.Itoa(weapons.plasma))
		}
		w.setWeaponDetails()
	}
	w.buildPlasma()
	berths.buildCrew()
	//.buildTotal()
}

func (w *weaponsDetail) particleChanged(value string) {
	particle, err := strconv.Atoi(value)
	if err == nil {
		weapons.accelerator = particle
		if w.countWeapons() > hull.maxHP {
			weapons.accelerator = particle - w.countWeapons() + hull.maxHP
			if weapons.accelerator < 0 {
				weapons.accelerator = 0
			}
			particleSelect.SetSelected(strconv.Itoa(weapons.accelerator))
		}
	}
	w.buildParticle()
	berths.buildCrew()
	// buildTotal()
}

func (w *weaponsDetail) setWeaponDetails() {
	w.buildMissile()
	w.buildBeam()
	w.buildPulse()
	w.buildPlasma()
	w.buildFusion()
	w.buildParticle()
}

func (w *weaponsDetail) buildMissile() {
	detailMissile.SetText(buildAmmoWeaponString("Triple Missile turrets: %d, tons: %d, ammo tons: %d", weapons.missile, int(float32(weapons.missile)+.9999), int(float32(4*weapons.missile)+.9999)))
	detailMissile.Refresh()
}

func (w *weaponsDetail) buildBeam() {
	detailBeam.SetText(buildWeaponString("Triple Beam laser turrets: %d, tons: %d", weapons.beam, int(float32(weapons.beam)+.9999)))
	detailBeam.Refresh()
}

func (w *weaponsDetail) buildPulse() {
	detailPulse.SetText(buildWeaponString("Triple Pulse lasr turrets: %d, tons: %d", weapons.pulse, int(float32(weapons.pulse)+.9999)))
	detailPulse.Refresh()
}

func (w *weaponsDetail) buildPlasma() {
	detailPlasma.SetText(buildWeaponString("Double Plasma gun turrets: %d, tons: %d", weapons.plasma, int(float32(2*weapons.plasma)+.9999)))
	detailPulse.Refresh()
}

func (w *weaponsDetail) buildFusion() {
	detailFusion.SetText(buildWeaponString("Double Fusion gun turrets: %d, tons: %d", weapons.fusion, int(float32(2*weapons.fusion)+.9999)))
	detailFusion.Refresh()
}

func (w *weaponsDetail) buildSand() {
	detailSand.SetText(buildAmmoWeaponString("Triple Sandcaster turrets: %d, tons: %d, ammo tons: %d", weapons.sandcaster, int(float32(weapons.sandcaster)/2.0+.9999), int(float32(weapons.sandcaster)+.9999)))
	detailFusion.Refresh()
}

func (w *weaponsDetail) buildParticle() {
	detailParticle.SetText(buildWeaponString("Particle Accelerator turrets: %d, tons: %d", weapons.accelerator, int(float32(3*weapons.accelerator)+.9999)))
	detailFusion.Refresh()
}

func (w *weaponsDetail) countWeapons() int {
	result := weapons.missile + weapons.beam + weapons.pulse + weapons.plasma + weapons.sandcaster + weapons.fusion + weapons.accelerator
	return result
}

func (w *weaponsDetail) buildWeapons() {
	w.buildMissile()
	w.buildBeam()
	w.buildPlasma()
	w.buildFusion()
	w.buildSand()
	w.buildParticle()
}

func (w *weaponsDetail) weaponsTonsUsed() int {
	result := int(.9999 + (5.0*float32(weapons.missile) + float32(weapons.beam) + float32(weapons.pulse) + 2.0*float32(weapons.fusion) + 2.0*float32(weapons.plasma) + 5.0*float32(weapons.accelerator)))
	return result
}

func buildWeaponString(weaponDescription string, count int, tons int) string {
	if count > 0 {
		return fmt.Sprintf(weaponDescription, count, int(float32(tons)))
	}
	return ""
}

func buildAmmoWeaponString(weaponAmmoDescription string, count int, tons int, ammoTons int) string {
	if count > 0 {
		return fmt.Sprintf(weaponAmmoDescription, count, tons, int(.999+float32(ammoTons)))
	}
	return ""
}

// Add next weapon, if needed, to the detailed list of weapons
func addWeapon(count int, box *widget.Box, label *widget.Label) {
	if count > 0 {
		weaponDetails.Children = append(weaponDetails.Children, label)
	}
}
