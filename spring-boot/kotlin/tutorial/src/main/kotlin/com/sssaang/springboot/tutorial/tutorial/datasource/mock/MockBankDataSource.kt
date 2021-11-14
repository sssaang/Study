package com.sssaang.springboot.tutorial.tutorial.datasource.mock

import com.sssaang.springboot.tutorial.tutorial.datasource.BankDataSource
import com.sssaang.springboot.tutorial.tutorial.model.Bank
import org.springframework.stereotype.Repository

@Repository
class MockBankDataSource: BankDataSource {
    private val banks: Collection<Bank> = listOf(
        Bank("123213", 0.1, 122),
        Bank("4251", 2.1, 12),
        Bank("1231251", 5.1, 12)
    )
    override fun retrieveBanks(): Collection<Bank> = banks
    override fun retrieveBank(accountNumber: String): Bank =
        banks.firstOrNull{ it.accountNumber == accountNumber }
            ?: throw NoSuchElementException("Could not found a bank with account number $accountNumber")
}