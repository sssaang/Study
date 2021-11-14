package com.sssaang.springboot.tutorial.tutorial.controller

import com.sssaang.springboot.tutorial.tutorial.model.Bank
import com.sssaang.springboot.tutorial.tutorial.service.BankService
import org.springframework.web.bind.annotation.GetMapping
import org.springframework.web.bind.annotation.RequestMapping
import org.springframework.web.bind.annotation.RestController

@RestController
@RequestMapping("api/banks")
class BankController(private val service: BankService) {

    @GetMapping
    fun getBanks(): Collection<Bank> = service.getBanks()
}