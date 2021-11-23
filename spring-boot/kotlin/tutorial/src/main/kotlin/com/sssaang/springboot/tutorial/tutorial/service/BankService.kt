package com.sssaang.springboot.tutorial.tutorial.service

import com.sssaang.springboot.tutorial.tutorial.datasource.BankDataSource
import com.sssaang.springboot.tutorial.tutorial.model.Bank
import org.springframework.stereotype.Service

@Service
class BankService(private val dataSource: BankDataSource) {
    fun getBanks(): Collection<Bank> = dataSource.retrieveBanks()
    fun getBank(accountNumber: String): Bank = dataSource.retrieveBank(accountNumber)
    fun addBank(bank: Bank): Bank = dataSource.createBank(bank)
}