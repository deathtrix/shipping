package main

import (
	"fmt"
	"log"
	"strconv"

	"github.com/deathtrix/shipping/packing"
	"github.com/gofiber/fiber/v3"
)

// Handler to get packages for the specified number of items
func getPackages(c fiber.Ctx) error {
	items, err := strconv.Atoi(c.Params("items"))
	if err != nil {
		log.Println(fmt.Errorf("getPackages: %w", err))
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Invalid packs"})
	}

	packs := packing.Calculate(packSizes, items)

	return c.Status(fiber.StatusOK).JSON(packs)
}

// Handler to get package sizes
func getPackageSizes(c fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(packSizes)
}

// Handler to set package sizes
func setPackageSizes(c fiber.Ctx) error {
	err := c.Bind().Body(&packSizes)
	if err != nil {
		log.Println(fmt.Errorf("setPackageSizes: %w", err))
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Invalid sizes"})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Sizes updated"})
}
