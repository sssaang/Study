package com.sssaang.springboot.tutorial.tutorial.datasource

import com.sssaang.springboot.tutorial.tutorial.model.Bank

interface BankDataSource {

    fun retrieveBanks(): Collection<Bank>
}

