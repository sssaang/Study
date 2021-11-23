package com.sssaang.springboot.tutorial.tutorial.datasource.mock

import com.sssaang.springboot.tutorial.tutorial.datasource.BankDataSource
import com.sssaang.springboot.tutorial.tutorial.model.Bank
import org.springframework.stereotype.Repository

@Repository
class MockBankDataSource: BankDataSource {
    private val banks: MutableList<Bank> = mutableListOf(
        Bank("123213", 0.1, 122),
        Bank("4251", 2.1, 12),
        Bank("1231251", 5.1, 12)
    )
    override fun retrieveBanks(): Collection<Bank> = banks
    override fun retrieveBank(accountNumber: String): Bank =
        banks.firstOrNull{ it.accountNumber == accountNumber }
            ?: throw NoSuchElementException("Could not found a bank with account number $accountNumber")

    override fun createBank(bank: Bank): Bank {
        if (banks.any { it.accountNumber == bank.accountNumber }) {
            throw IllegalArgumentException("Bank with account number ${bank.accountNumber} already exists")
        }
        banks.add(bank)
        return bank
    }

    override fun updateBank(bank: Bank): Bank {
        val updatingBank = banks.firstOrNull{ it.accountNumber == bank.accountNumber }
            ?: throw NoSuchElementException("Could not found a bank with account number ${bank.accountNumber}")

        banks.remove(updatingBank)
        banks.add(bank)
        return bank
    }

}