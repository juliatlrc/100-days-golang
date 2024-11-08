package annalyn

// CanFastAttack can be executed only when the knight is sleeping.
func CanFastAttac(knightIsAwake bool) bool {
    return !knightIsAwake // Retorna true se o cavaleiro estiver dormindo
}

// CanSpy can be executed if at least one of the characters is awake.
func CanSp(knightIsAwake, archerIsAwake, prisonerIsAwake bool) bool {
    return knightIsAwake || archerIsAwake || prisonerIsAwake // Retorna true se qualquer um estiver acordado
}

// CanSignalPrisoner can be executed if the prisoner is awake and the archer is sleeping.
func CanSignalPrisone(archerIsAwake, prisonerIsAwake bool) bool {
    return !archerIsAwake && prisonerIsAwake // Retorna true se o arqueiro estiver dormindo e o prisioneiro acordado
}

// CanFreePrisoner can be executed if the prisoner is awake and the other 2 characters are asleep
// or if Annalyn's pet dog is with her and the archer is sleeping.
func CanFreePrisone(knightIsAwake, archerIsAwake, prisonerIsAwake, petDogIsPresent bool) bool {
    // Se o cachorro está presente, e o arqueiro está dormindo, podemos tentar libertar o prisioneiro
    if petDogIsPresent {
        return !archerIsAwake // Libertação se o arqueiro estiver dormindo e o cachorro estiver presente
    }

    // Se o cachorro não está presente, verificamos a situação dos outros dois
    return !knightIsAwake && !archerIsAwake && prisonerIsAwake // Libertação se ambos os guardas estão dormindo e o prisioneiro está acordado
}

// o CanFreePrisone não vai funcionar, pq tem alguns casos tipo, que o prisioneiro e cavaleiros acordados e o cachorro presente, ele vai conseguir escapar 
