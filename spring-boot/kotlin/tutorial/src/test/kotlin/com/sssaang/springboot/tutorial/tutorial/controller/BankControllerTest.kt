package com.sssaang.springboot.tutorial.tutorial.controller

import com.fasterxml.jackson.databind.ObjectMapper
import com.sssaang.springboot.tutorial.tutorial.model.Bank
import org.junit.jupiter.api.Assertions.*
import org.junit.jupiter.api.DisplayName
import org.junit.jupiter.api.Nested
import org.junit.jupiter.api.Test
import org.junit.jupiter.api.TestInstance
import org.springframework.beans.factory.annotation.Autowired
import org.springframework.boot.test.autoconfigure.web.servlet.AutoConfigureMockMvc
import org.springframework.boot.test.context.SpringBootTest
import org.springframework.http.MediaType
import org.springframework.test.web.servlet.MockMvc
import org.springframework.test.web.servlet.get
import org.springframework.test.web.servlet.patch
import org.springframework.test.web.servlet.post

@SpringBootTest
@AutoConfigureMockMvc
internal class BankControllerTest {

    @Autowired
    lateinit var mockMvc: MockMvc

    @Autowired
    lateinit var objectMapper: ObjectMapper

    @Nested
    @DisplayName("GET /api/bank")
    @TestInstance(TestInstance.Lifecycle.PER_CLASS)
    inner class GetBanks {
        @Test
        fun `should return all banks`() {
            mockMvc.get("/api/banks")
                .andDo { print() }
                .andExpect {
                    status { isOk() }
                    content { contentType((MediaType.APPLICATION_JSON)) }
                }
        }
    }

    @Nested
    @DisplayName("GET /api/bank/{accountNumber}")
    @TestInstance(TestInstance.Lifecycle.PER_CLASS)
    inner class GetBank {
        @Test
        fun `should return the bank with given account number`() {
            val accountNumber = "123213"

            mockMvc.get("/api/banks/$accountNumber")
                .andDo { print() }
                .andExpect {
                    status { isOk() }
                    content { contentType((MediaType.APPLICATION_JSON)) }
                    jsonPath("$.trust") { value("0.1") }
                    jsonPath("$.transactionFee") { value("122") }
                }
        }

        @Test
        fun `should return NOT FOUND if the account number does not exist`() {
            val accountNumber = "does_not_exist"

            mockMvc.get("/api/banks/${accountNumber}")
                .andDo { print() }
                .andExpect { status { isNotFound() } }
        }
    }

    @Nested
    @DisplayName("POST /api/banks")
    @TestInstance(TestInstance.Lifecycle.PER_CLASS)
    inner class PostBank {
        @Test
        fun `should add the new bank`() {
            val newBank = Bank("acc123", 31.415, 2)

            val performPost = mockMvc.post("/api/banks") {
                contentType = MediaType.APPLICATION_JSON
                content = objectMapper.writeValueAsString(newBank)
            }

            performPost.andDo { print() }
                .andExpect {
                    status { isCreated() }
                    content { contentType(MediaType.APPLICATION_JSON) }
                    jsonPath("$.accountNumber") { value("acc123") }
                    jsonPath("$.trust") { value(31.415) }
                    jsonPath("$.transactionFee") { value(2) }
                }

            mockMvc.get("/api/banks/${newBank.accountNumber}")
                .andDo { print() }
                .andExpect {
                    status { isOk() }
                    content {
                        json(objectMapper.writeValueAsString(newBank))
                    }
                }
        }

        @Test
        fun `should return BAD REQUEST if bank with given account number already exists`() {
            val invalidBank = Bank("123213", 0.1, 122)

            val performPost = mockMvc.post("/api/banks") {
                contentType = MediaType.APPLICATION_JSON
                content = objectMapper.writeValueAsString(invalidBank)
            }

            performPost.andDo { print() }
                .andExpect {
                    status { isBadRequest() }
                }
        }
    }

    @Nested
    @DisplayName("PATCH /api/banks")
    @TestInstance(TestInstance.Lifecycle.PER_CLASS)
    inner class PatchExistingBank {
        @Test
        fun `should update an existing bank`() {
            val updatedBank = Bank("123213", 1.1, 11)

            val performPatchRequest = mockMvc.patch("/api/banks") {
                contentType = MediaType.APPLICATION_JSON
                content = objectMapper.writeValueAsString(updatedBank)
            }

            performPatchRequest
                .andDo { print() }
                .andExpect {
                    status { isOk() }
                    content {
                        contentType(MediaType.APPLICATION_JSON )
                        json(objectMapper.writeValueAsString(updatedBank))
                    }
                }

            mockMvc.get("/api/banks/${updatedBank.accountNumber}")
                .andDo { print() }
                .andExpect {
                    status { isOk() }
                    content {
                        json(objectMapper.writeValueAsString(updatedBank))
                    }
                }
        }

        @Test
        fun `should return BAD REQUEST if bank with given account number does not exists`() {
            val invalidBank = Bank("12", 0.1, 122)

            val performPost = mockMvc.post("/api/banks") {
                contentType = MediaType.APPLICATION_JSON
                content = objectMapper.writeValueAsString(invalidBank)
            }

            performPost.andDo { print() }
                .andExpect {
                    status { isNotFound() }
                }
        }
    }
}